package config

import (
	"flag"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Hooks   []Hook  `yaml:"hooks" env-required:"true"`
	Logging Logging `yaml:"logging"`
}

type Hook struct {
	Name    string   `yaml:"name" env-required:"true"`
	Keys    []string `yaml:"keys" env-required:"true"`
	Command string   `yaml:"command" env-required:"true"`
	Args    []string `yaml:"args"`
}

type Logging struct {
	Level  string `yaml:"level" env-default:"info"`
	Output string `yaml:"output" env-default:"stdout"`
}

func MustRead(path string) (cfg Config) {
	const operation = "config.MustReadConfig"

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(fmt.Errorf("%s: %w", operation, err))
	}

	return cfg
}

func MustReadFromFlag() Config {
	const operation = "config.MustReadFromFlag"
	var path string

	flag.StringVar(&path, "path", "", "path to config file")
	flag.Parse()

	if path == "" {
		panic(fmt.Errorf("%s: path is required", operation))
	}

	return MustRead(path)
}
