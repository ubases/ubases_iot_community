package iotmqtt

import (
	"context"
	"crypto/tls"
	"errors"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

type ConnHandler interface {
	OnConnectionLostHandler(err error)
	OnConnectHandler()
	GetOnlineMsg() (string, string)
	GetWillMsg() (string, string)
}

type Client struct {
	Options     ClientOptions
	client      paho.Client
	router      *router
	connHandler ConnHandler
}

type ClientOptions struct {
	Servers       []string
	ClientID      string
	Username      string
	Password      string
	TLSConfig     *tls.Config
	AutoReconnect bool
	Keepalive     int
	Clean         bool
}

type QOS byte

const (
	AtMostOnce QOS = iota
	AtLeastOnce
	ExactlyOnce
)

var (
	ErrMinimumOneServer = errors.New("mqtt: at least one server needs to be specified")
)

func handle(callback MessageHandler) paho.MessageHandler {
	return func(client paho.Client, message paho.Message) {
		if callback != nil {
			callback(Message{message: message})
		}
	}
}

func NewClient(options ClientOptions, connhandler ConnHandler) (*Client, error) {
	pahoOptions := paho.NewClientOptions()
	if options.Servers != nil && len(options.Servers) > 0 {
		for _, server := range options.Servers {
			pahoOptions.AddBroker(server)
		}
	} else {
		return nil, ErrMinimumOneServer
	}

	if options.ClientID == "" {
		options.ClientID = uuid.New().String()
	}
	pahoOptions.SetClientID(options.ClientID)
	pahoOptions.SetProtocolVersion(4)
	pahoOptions.SetKeepAlive(time.Duration(options.Keepalive) * time.Second)
	pahoOptions.SetWriteTimeout(10 * time.Second)
	pahoOptions.SetConnectTimeout(10 * time.Second)
	pahoOptions.SetPingTimeout(10 * time.Second)

	if options.TLSConfig != nil {
		pahoOptions.SetTLSConfig(options.TLSConfig)
	}

	if options.Username != "" {
		pahoOptions.SetUsername(options.Username)
		pahoOptions.SetPassword(options.Password)
	}

	pahoOptions.SetAutoReconnect(options.AutoReconnect)
	pahoOptions.SetMaxReconnectInterval(15 * time.Second)

	pahoOptions.SetCleanSession(options.Clean)

	var client Client
	client.connHandler = connhandler
	if connhandler != nil {
		pahoOptions.SetConnectionLostHandler(client.ConnectionLostHandler)
		pahoOptions.SetOnConnectHandler(client.OnConnectHandler)
		if t, m := connhandler.GetWillMsg(); t != "" {
			pahoOptions.SetWill(t, m, 0, true)
		}
	}

	pahoClient := paho.NewClient(pahoOptions)
	router := newRouter()
	pahoClient.AddRoute("#", handle(func(message Message) {
		routes := router.match(&message)
		for _, route := range routes {
			m := message
			m.vars = route.vars(&message)
			route.handler(m)
		}
	}))

	client.client = pahoClient
	client.Options = options
	client.router = router

	return &client, nil
}

func (c *Client) Connect(ctx context.Context) error {
	token := c.client.Connect()
	return tokenWithContext(ctx, token)
}

func (c *Client) IsConnected() bool {
	return c.client.IsConnected()
}

func (c *Client) IsConnectionOpen() bool {
	return c.client.IsConnectionOpen()
}

func (c *Client) DisconnectImmediately() {
	c.client.Disconnect(3000)
}

/*
func tokenWithContext(ctx context.Context, token paho.Token) error {
	completer := make(chan error)

	go func() {
		token.Wait()
		completer <- token.Error()
	}()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-completer:
			return err
		}
	}
}
*/
func tokenWithContext(ctx context.Context, token paho.Token) error {
	completer := make(chan error)
	go func() {
		var err error
		if ret := token.WaitTimeout(60 * time.Second); ret {
			err = token.Error()
		} else {
			err = errors.New("timeout")
		}
		completer <- err
	}()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-completer:
			return err
		}
	}
}

func (c *Client) ConnectionLostHandler(client paho.Client, err error) {
	if c.connHandler != nil {
		c.connHandler.OnConnectionLostHandler(err)
	}
}

func (c *Client) OnConnectHandler(client paho.Client) {
	if c.connHandler != nil {
		c.connHandler.OnConnectHandler()
		//上线消息
		topic, payload := c.connHandler.GetOnlineMsg()
		if topic != "" {
			_ = c.PublishString(context.Background(), topic, payload, AtLeastOnce, Retain)
		}
	}
}
