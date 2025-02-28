package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env  string     `yaml:"env" env-default:"local"`
	SMTP SMTPConfig `yaml:"smtp"`
}

var cfgInstance *Config

type SMTPConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-default:"587"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
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
		panic("config file not found " + configPath)
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
