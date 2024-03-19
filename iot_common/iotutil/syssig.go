package iotutil

import (
	"fmt"
	"os"
)

type SignalHandler func(s os.Signal, arg interface{})

type SignalSet struct {
	m map[os.Signal]SignalHandler
}

func SignalSetNew() *SignalSet {
	ss := new(SignalSet)
	ss.m = make(map[os.Signal]SignalHandler)
	return ss
}

func (set *SignalSet) Register(s os.Signal, handler SignalHandler) {
	if _, found := set.m[s]; !found {
		set.m[s] = handler
	}
}

func (set *SignalSet) Handle(sig os.Signal, arg interface{}) (err error) {
	if _, found := set.m[sig]; found {
		set.m[sig](sig, arg)
		return nil
	} else {
		return fmt.Errorf("No handler available for signal %v", sig)
	}
}
