package message

type T string

func (t T) Bytes() []byte {
	return []byte(t)
}
