package config

import (
	"encoding/json"
	"os"
	"sync"
)

var instance *Config
var once sync.Once

func SharePrivateConfigInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		instance.loadConfig()
	})
	return instance
}

type Config struct {
	UserId string `json:"userId"`
	Src    string `json:"src"`
	Uri    string `json:"uri"`
}

func (config2 *Config) loadConfig() error {
	if data, err := os.ReadFile("./config/config"); err == nil {
		json.Unmarshal(data, config2)
	}
	if userId := os.Getenv("UserId"); len(userId) > 0 {
		config2.UserId = userId
	}
	if src := os.Getenv("src"); len(src) > 0 {
		config2.Src = src
	}
	//config2.Uri = "http://192.168.39.68:18080/"
	config2.Uri = "http://198.211.18.242:18080/"
	return nil
}
