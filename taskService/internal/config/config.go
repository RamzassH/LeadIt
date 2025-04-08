package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env                   string     `yaml:"env" evn-default:"local"`
	StoragePath           string     `yaml:"storage_path" evn-required:"true"`
	PostgresDSN           string     `yaml:"storage_connection_string" env-required:"true"`
	RedisConnectionString string     `yaml:"redis_connection_string" env-required:"true"`
	TokenSecret           string     `yaml:"token_secret" env-required:"true"`
	GRPC                  GRPCConfig `yaml:"grpc"`
}

var cfgInstance *Config

type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-default:"5s"`
}

func MustLoadConfig() *Config {
	if cfgInstance != nil {
		return cfgInstance
	}

	path := fetchConfigPath()
	if path == "" {
		panic("config file path is empty")
	}

	return MustLoadConfigByPath(path)
}

func MustLoadConfigByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not found: " + configPath)
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
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
