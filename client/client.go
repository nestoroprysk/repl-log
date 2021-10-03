package client

import (
	"fmt"
	"net/http"
	"time"

	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/message"

	"github.com/go-resty/resty/v2"
)

type T struct {
	*resty.Client
	address string
}

func New(c config.T) (*T, error) {
	result := &T{
		Client:  resty.New(),
		address: "http://" + c.Address(),
	}

	if err := result.Ping(); err != nil {
		return nil, err
	}

	return result, nil
}

func (t *T) Ping() error {
	var m message.T
	resp, err := t.SetRetryCount(10).
		SetRetryWaitTime(1 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second).
		R().SetResult(&m).Get(t.address + "/ping")
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("expecting status OK, got: %s", resp.Status)
	}
	if m != message.T("pong") {
		return fmt.Errorf("expecting pong, got: %s", m)
	}

	return nil
}

func (t *T) GetMessages() ([]message.T, error) {
	var result []message.T
	resp, err := t.R().SetResult(&result).Get(t.address + "/messages")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("expecting status OK, got: %s", resp.StatusCode())
	}

	return result, nil
}

func (t *T) PostMessage(m message.T) error {
	var result message.T
	resp, err := t.R().SetBody(m.Bytes()).SetResult(&result).Post(t.address + "/messages")
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf("expecting status created, got: %s", resp.StatusCode())
	}
	if result != m {
		return fmt.Errorf("expecting the sent message %v, got: %v", m, resp)
	}

	return nil
}
