package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Library []*PathConfig `json:"library"`

	ServerInfo struct {
		Ip   string
		Port string
	}
}

type PathConfig struct {
	Path   string `json:"path"`
	Prefix string `json:"prefix"`
}

func NewConfig(file string) *Config {
	c := new(Config)

	if file, err := os.Open(file); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
