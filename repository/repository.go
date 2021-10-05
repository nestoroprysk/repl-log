package repository

import (
	"sync"

	"github.com/nestoroprysk/repl-log/message"
)

type T struct {
	sync.RWMutex
	messages map[message.Namespace][]message.T
}

func New() *T {
	result := &T{messages: map[message.Namespace][]message.T{}}
	return result
}

func (t *T) GetMessages(ns ...message.Namespace) []message.T {
	t.RLock()
	defer t.RUnlock()

	if len(ns) == 0 {
		return t.messages[message.DefaultNamespace]
	}

	var result []message.T
	for _, n := range ns {
		ms := t.messages[n]
		result = append(result, ms...)
	}

	return result
}

func (t *T) AppendMessage(m message.T) message.T {
	t.Lock()
	defer t.Unlock()

	if m.Namespace == "" {
		m.Namespace = message.DefaultNamespace
	}

	t.messages[m.Namespace] = append(t.messages[m.Namespace], m)
	return m
}
