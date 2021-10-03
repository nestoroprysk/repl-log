package config

type T struct {
	Host string
	Port string
}

func (t T) Address() string {
	return t.Host + ":" + t.Port
}
