package client

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/nestoroprysk/repl-log/config"
	"github.com/nestoroprysk/repl-log/message"

	"github.com/go-resty/resty/v2"
)

type T struct {
	*resty.Client
	address string
}

func New(c config.Location) (*T, error) {
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
	resp, err := t.R().Get(t.address + "/ping")
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("expecting status OK, got: %s", resp.Status)
	}
	if resp.String() != "pong" {
		return fmt.Errorf("expecting pong, got: %s", resp.String())
	}

	return nil
}

func (t *T) GetNamespaces() ([]message.Namespace, error) {
	var result []message.Namespace
	resp, err := t.R().SetResult(&result).Get(t.address + "/namespaces")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("expecting status OK, got: %s", resp.StatusCode())
	}

	return result, nil
}

func (t *T) GetMessages(ns ...message.Namespace) ([]message.T, error) {
	var result []message.T
	resp, err := t.R().SetResult(&result).
		SetQueryParamsFromValues(url.Values{
			"namespace": toStrings(ns...),
		}).Get(t.address + "/messages")
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
	resp, err := t.R().SetBody(m).SetResult(&result).Post(t.address + "/messages")
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

func (t *T) DeleteNamespace(n message.Namespace) (bool, error) {
	resp, err := t.R().Delete(t.address + fmt.Sprintf("/namespaces/%s", n))
	if err != nil {
		return false, err
	}

	if resp.StatusCode() == http.StatusOK {
		return true, nil
	}
	if resp.StatusCode() == http.StatusNoContent {
		return false, nil
	}

	return false, fmt.Errorf("unexpected status code %s", resp.StatusCode())
}

func toStrings(ns ...message.Namespace) []string {
	var result []string
	for _, n := range ns {
		result = append(result, string(n))
	}

	return result
}
