package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	HealthCheckInterval string   `json:"healthCheckInterval"`
	Servers             []string `json:"servers"`
	ListenPort          string   `json:"listenPort"`
}

func LoadConfig(file string) (Config, error) {
	var config Config

	bytes, err := os.ReadFile(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
