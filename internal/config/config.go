package config

import (
	"PhoneBook/pkg/logger"
	"PhoneBook/pkg/token"
)

type Config struct {
	Logger *logger.Config `koanf:"logger"`
	Token  *token.Config  `koanf:"token"`
}
