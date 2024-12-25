package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env             string        `yaml:"env" env-default:"local"`
	StoragePath     string        `yaml:"storage_path" env-required:"true"`
	TokenTTL        time.Duration `yaml:"token_ttl" env-required:"true"`
	TokenSecret     string        `yaml:"token_secret" env-required:"true"`
	RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env-required:"true"`
	PostgresDSN     string        `yaml:"storage_connection_string" env-required:"true"`
	GRPC            GRPCConfig    `yaml:"grpc"`
}

var cfgInstance *Config

type GRPCConfig struct {
	Port    int    `yaml:"port" env-default:"8080"`
	Timeout string `yaml:"timeout" env-default:"5s"`
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

	fmt.Println("Using config path:", res)

	return res
}
