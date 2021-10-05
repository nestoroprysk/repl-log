package config

import (
	"os"

	"github.com/nestoroprysk/repl-log/env"
)

type T struct {
	Host string
	Port string
}

var (
	Master     = T{Host: os.Getenv(env.MasterHost), Port: os.Getenv(env.MasterPort)}
	Secondary  = T{Host: os.Getenv(env.SecondaryHost), Port: os.Getenv(env.SecondaryPort)}
	SecondaryA = T{Host: os.Getenv(env.SecondaryAHost), Port: os.Getenv(env.SecondaryAPort)}
	SecondaryB = T{Host: os.Getenv(env.SecondaryBHost), Port: os.Getenv(env.SecondaryBPort)}
)

func (t T) Address() string {
	return t.Host + ":" + t.Port
}
