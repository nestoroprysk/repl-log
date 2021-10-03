package repository

import (
	"sync"

	"github.com/nestoroprysk/repl-log/message"
)

type T struct {
	sync.RWMutex
	messages []message.T
}

func New() *T {
	result := &T{}
	return result
}

func (t *T) GetMessages() []message.T {
	t.RLock()
	defer t.RUnlock()
	return t.messages
}

func (t *T) AppendMessage(m message.T) message.T {
	t.Lock()
	defer t.Unlock()
	t.messages = append(t.messages, m)
	return m
}
