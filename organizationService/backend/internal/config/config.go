package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env         string    `yaml:"env" env-default:"local"`
	StoragePath string    `yaml:"storage_path" env-required:"true"`
	PostgresDSN string    `yaml:"storage_connection_string" env-required:"true"`
	GRPC        GRPConfig `yaml:"grpc"`
}

var cfgInstance *Config

type GRPConfig struct {
	Port    int `yaml:"port" env-default:"8080"`
	Timeout int `yaml:"timeout" env-default:"5s"`
}

func MustLoadConfig() *Config {
	if cfgInstance != nil {
		return cfgInstance
	}

	path := fetchConfigPath()
	if path == "" {
		panic("config file path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found " + path)
	}

	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}

	cfgInstance = &config
	return cfgInstance
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "./local.yaml", "config file path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
