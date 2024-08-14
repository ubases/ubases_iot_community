package common

import (
	"sync"
)

type EventData struct {
	Subject string
	Data    string
}

type Subscriber struct {
	ID       string
	Subjects map[string]bool
	Channel  chan EventData
}

type Publisher struct {
	SubDict map[string]*Subscriber
	mutex   *sync.Mutex
}

func NewPublisher() *Publisher {
	return &Publisher{
		SubDict: make(map[string]*Subscriber),
		mutex:   &sync.Mutex{},
	}
}

func (p *Publisher) AddSubscriber(s *Subscriber) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.SubDict[s.ID] = s
}

func (p *Publisher) DelSubscriber(s *Subscriber) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.SubDict, s.ID)
}

func (p *Publisher) Publish(ed EventData) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	wg := &sync.WaitGroup{}
	for _, s := range p.SubDict {
		//没有订阅该主题的订阅者,不推送消息
		if _, ok := s.Subjects[ed.Subject]; !ok {
			continue
		}
		wg.Add(1)
		go func(sub *Subscriber) {
			defer wg.Done()
			sub.Channel <- ed
		}(s)
	}
	wg.Wait()
}
