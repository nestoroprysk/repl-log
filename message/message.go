package message

import "sync/atomic"

type T struct {
	Message   string `json:"message"`
	Namespace `json:"namespace"`
	ID        uint32 `json:"id"`
	// Delay in milliseconds.
	Delay uint32 `json:"delay"`
}

type Namespace string

var DefaultNamespace Namespace = "default"

const NamespaceID = "id"

var id uint32

// NextID returns a bigger value on each call.
func NextID() uint32 {
	return atomic.AddUint32(&id, 1)
}
