package config

import (
	"gw-exchanger/pkg/logger"
	"gw-exchanger/pkg/utils"
)

var Cfg Config

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     uint16
}

func NewConfig() Config {
	logger.Log.Debug("Loading config")

	config := Config{
		DBUser:     utils.GetEnv("POSTGRES_USER", DefaultDBUser),
		DBPassword: utils.GetEnv("POSTGRES_PASSWORD", DefaultDBPassword),
		DBName:     utils.GetEnv("POSTGRES_DB", DefaultDBName),
		DBHost:     utils.GetEnv("POSTGRES_SERVER", DefaultDBHost),
		DBPort:     utils.GetEnvUint16("POSTGRES_PORT", DefaultDBPort),
	}

	logger.Log.Debug("Loaded config", "config", config)

	return config
}
