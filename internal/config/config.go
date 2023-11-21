package config

import "PhoneBook/pkg/logger"

type Config struct {
	Logger *logger.Config `koanf:"logger"`
}
