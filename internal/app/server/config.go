package server

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Addr     string `json:"addr"`
	DBPath   string `json:"db_path"`
	DBDriver string `json:"db_driver"`
}

func NewConfig() *Config {
	return &Config{
		Addr: ":17555",
	}
}

func ReadConfig(path string, config *Config) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, config); err != nil {
		return err
	}

	return nil
}
