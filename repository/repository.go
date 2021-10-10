package repository

import (
	"sort"
	"sync"

	"github.com/nestoroprysk/repl-log/message"
)

type T struct {
	sync.RWMutex
	messages   map[message.Namespace][]message.T
	namespaces map[message.Namespace]bool
}

func New() *T {
	return &T{
		messages: map[message.Namespace][]message.T{},
		namespaces: map[message.Namespace]bool{
			message.DefaultNamespace: true,
		},
	}
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

func (t *T) GetNamespaces() []message.Namespace {
	t.RLock()
	defer t.RUnlock()

	var result []message.Namespace
	for n := range t.namespaces {
		result = append(result, n)
	}

	return result
}

func (t *T) DeleteNamespace(n message.Namespace) bool {
	t.Lock()
	defer t.Unlock()

	if t.namespaces[n] == false {
		return false
	}

	delete(t.namespaces, n)
	delete(t.messages, n)

	return true
}

func (t *T) AppendMessage(m message.T) message.T {
	t.Lock()
	defer t.Unlock()

	if m.Namespace == "" {
		m.Namespace = message.DefaultNamespace
	}

	t.namespaces[m.Namespace] = true
	t.messages[m.Namespace] = insertSorted(t.messages[m.Namespace], m)
	return m
}

func insertSorted(ms []message.T, m message.T) []message.T {
	i := sort.Search(len(ms), func(i int) bool { return ms[i].ID >= m.ID })
	ms = append(ms, message.T{})
	copy(ms[i+1:], ms[i:])
	ms[i] = m
	return ms
}
