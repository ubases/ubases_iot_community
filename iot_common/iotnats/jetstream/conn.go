package jetstream

import (
	"errors"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

type JetStreamIface interface {
	GetJetStreamContext() nats.JetStreamContext
	Close()
}

type JetStream struct {
	config    *Config
	nc        *nats.Conn
	jsContext nats.JetStreamContext
}

func (js *JetStream) Close() {
	_ = js.nc.FlushTimeout(5 * time.Second)
	js.nc.Close()
}

func (js *JetStream) GetJetStreamContext() nats.JetStreamContext {
	return js.jsContext
}

func expandClusterURL(href string) string {
	if strings.Contains(href, ",") {
		return href
	}

	u, err := url.Parse(href)
	if err != nil {
		return href
	}
	addrs, err := net.LookupHost(u.Hostname())
	if err != nil {
		return href
	}
	if len(addrs) == 0 {
		return href
	}

	cl := ""
	for i, nm := range addrs {
		if i > 0 {
			cl += ","
		}
		cl += u.Scheme + "://"
		if u.User != nil {
			cl += u.User.String()
			cl += "@"
		}
		cl += nm
		if u.Port() != "" {
			cl += ":" + u.Port()
		}
		if u.RawQuery != "" {
			cl += "?" + u.RawQuery
		}
	}

	return cl
}

func NewJetStream(cfg *Config) (JetStreamIface, error) {
	urls := expandClusterURL(cfg.ClusterURLs)
	nc, err := nats.Connect(urls,
		nats.Name(cfg.Name),
		nats.RetryOnFailedConnect(true),
		//nats.MaxReconnects(36),//不限制重连次数
		nats.ReconnectWait(5*time.Second),
		nats.ReconnectHandler(reconnectHandler),
		nats.DisconnectErrHandler(cfg.ConnErrHandler),
	)

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	j := &JetStream{
		config:    cfg,
		nc:        nc,
		jsContext: js,
	}

	if err := j.createStream(); err != nil {
		nc.Close()
		return nil, err
	}

	return j, nil
}

func (js *JetStream) createStream() error {
	if js.config.StreamConfig == nil {
		return errors.New("StreamConfig is required")
	}
	if js.config.StreamConfig.Name == "" {
		return errors.New("StreamConfig.Name is required")
	}
	subjs := js.config.StreamConfig.Subjects
	if subjs == nil {
		subjs = []string{js.config.StreamConfig.Name + ".*"}
		js.config.StreamConfig.Subjects = subjs
	}

	var stream *nats.StreamInfo
	var err error
	stream, err = js.jsContext.StreamInfo(js.config.StreamConfig.Name)
	if err != nil {
	}
	if stream == nil {
		stream, err = js.jsContext.AddStream(js.config.StreamConfig)
		if err != nil {
			return err
		}
	}
	return nil
}
