package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Fields `yaml:"fields"`
}

type Fields struct {
	Level      string `yaml:"level" env:"LEVEL"`
	Message    string `yaml:"message" env:"MESSAGE"`
	Time       string `yaml:"time" env:"TIME"`
	TimeLayout string `yaml:"timeLayout" env:"TIMELAYOUT"`
}

func MustLoad(path string) Config {
	if path == "" {
		log.Fatal("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("error reading config: %s", path)
	}

	return cfg
}
