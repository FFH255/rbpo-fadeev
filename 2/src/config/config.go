package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Alg     string `yaml:"alg" default:"md5"`
	Length  int    `yaml:"length" default:"5"`
	Chars   string `yaml:"chars" default:"abcdefghijklmnopqrstuvwxyz"`
	Workers int    `yaml:"workers" default:"12"`
}

func MustLoad() *Config {
	configPath := flag.String("config", "", "path to yaml configure file.")

	flag.Parse()

	if *configPath == "" {
		log.Fatalf("config command flag is not set")
	}

	if _, err := os.Stat(*configPath); err != nil {
		log.Fatalf("error opening config file: %s", err)
	}

	config := new(Config)

	if err := cleanenv.ReadConfig(*configPath, config); err != nil {
		log.Fatalf("error reading config file: %s", err)
	}

	return config
}
