package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Location struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Listen    Location   `json:"listen"`
	Replicate []Location `json:"replicate"`
}

const (
	configPath = "CONFIG_PATH"
)

func (t Location) Address() string {
	return t.Host + ":" + strconv.Itoa(t.Port)
}

func Read() (Config, error) {
	configPath := os.Getenv(configPath)
	if configPath == "" {
		return Config{}, fmt.Errorf("set the environmental variable %s", configPath)
	}

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}

	var result Config
	if err := json.Unmarshal(b, &result); err != nil {
		return Config{}, err
	}

	return result, nil
}
