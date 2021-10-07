package util

import (
	"github.com/nestoroprysk/repl-log/client"
	"github.com/nestoroprysk/repl-log/config"
)

func ToClients(locations []config.Location) ([]*client.T, error) {
	var result []*client.T
	for _, l := range locations {
		c, err := client.New(l)
		if err != nil {
			return nil, err
		}

		result = append(result, c)
	}

	return result, nil
}
