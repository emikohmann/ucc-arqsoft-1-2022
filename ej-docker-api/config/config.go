package config

import (
	"fmt"
	"github.com/emikohmann/ucc-arqsoft-1/ej-docker-api/utils/strs"
	"os"
	"strings"
)

const (
	envPort        = "PORT"
	envEnvironment = "ENVIRONMENT"
)

func Port() string {
	port := os.Getenv(envPort)
	if strs.IsEmpty(port) {
		return ":8080" // default port
	}
	if !strings.HasPrefix(port, ":") {
		return fmt.Sprintf(":%s", port)
	}
	return port
}

func IsTest() bool {
	return os.Getenv(envEnvironment) == "test"
}
