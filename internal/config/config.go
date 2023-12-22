package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	Port int    `yaml:"port"`

	WriteTimeout time.Duration `yaml:"writeTimeout"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	f, err := os.ReadFile(path)

	var cfg Config
	if err = yaml.Unmarshal(f, &cfg); err != nil {
		panic("failed to unmarshal config file: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string
	flag.StringVar(&path, "c", "config/*.yaml", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
