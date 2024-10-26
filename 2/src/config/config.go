package config

import (
	"flag"
	"log"
	"os"
	"ssd-lab-pswd-go/src/bruteforce"
	"ssd-lab-pswd-go/src/generator"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HashGoal   string            `yaml:"hash_goal"`
	Generator  generator.Config  `yaml:"generator"`
	Bruteforce bruteforce.Config `yaml:"bruteforce"`
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
