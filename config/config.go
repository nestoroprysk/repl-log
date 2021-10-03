package config

import (
	"fmt"
	"os"
)

type T struct {
	host string
	port string
}

const (
	envHost = "HOST"
	envPort = "PORT"
)

func Make() (T, error) {
	host := os.Getenv(envHost)
	if host == "" {
		return T{}, fmt.Errorf("initialize the environmental variable %q", envHost)
	}

	port := os.Getenv(envPort)
	if port == "" {
		return T{}, fmt.Errorf("initialize the environmental variable %q", envPort)
	}

	return T{
		host: host,
		port: port,
	}, nil
}

func (t T) Address() string {
	return t.host + ":" + t.port
}
