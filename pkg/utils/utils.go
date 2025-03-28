package utils

import (
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetEnvUint16(key string, fallback uint16) uint16 {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	res, err := strconv.ParseUint(value, 10, 16)
	if err != nil {
		return fallback
	}

	return uint16(res)
}
