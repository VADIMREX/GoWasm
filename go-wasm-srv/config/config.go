package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port int16
}

func ReadConfig(path string) Config {
	cfg := Config{Port: 8080}
	raw, err := os.ReadFile(path)
	if err != nil {
		return cfg
	}
	json.Unmarshal(raw, &cfg)
	return cfg
}
