package server

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port     string `json:"port"`
	DbPath   string `json:"db_path"`
	DbDriver string `json:"db_driver"`
}

func NewConfig() *Config {
	return &Config{
		Port: ":6970",
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
