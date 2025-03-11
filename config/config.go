package config

import (
	"os"
)

// GetPort returns the server port
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
