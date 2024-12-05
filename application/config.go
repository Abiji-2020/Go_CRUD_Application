package application

import (
	"os"
	"strconv"
)

type config struct {
	RedisAddress string
	ServerPort   uint16
}

func LoadConfig() config {
	cfg := config{
		RedisAddress: "localhost:6379",
		ServerPort:   3000,
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		cfg.RedisAddress = redisAddr
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		} else {
			// Handle error if port cannot be parsed
			// For now, just log the error
			println("Error parsing SERVER_PORT:", err)
		}
	}

	return cfg
}