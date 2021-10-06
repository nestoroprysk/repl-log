package message

type T struct {
	Message   string `json:"message"`
	Namespace `json:"namespace"`
}

type Namespace string

var DefaultNamespace Namespace = "default"

const NamespaceID = "id"
