package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoadConfig()

	log.Info().Msg("Starting notification service")

	validate := validator.New()

	logger := set
}
