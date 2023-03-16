package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Configuration defintion.
type Config struct {
	LogLevel string
	Server   *ServerConfig
}

// ServerConfig declaration.
type ServerConfig struct {
	ServerPort   int
	PprofEnabled bool
}

// GetServerConfig call.
func GetServerConfig() (*ServerConfig, error) {
	serverPort, err := getInt("SERVER_PORT")
	if err != nil {
		return nil, err
	}

	pprofEnabled, err := getBool("PPROF_ENABLED")
	if err != nil {
		return nil, err
	}

	return &ServerConfig{
		ServerPort:   serverPort,
		PprofEnabled: pprofEnabled,
	}, nil
}

func LoadConfiguration() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	logLevel, err := get("LOG_LEVEL")
	if err != nil {
		return nil, err
	}

	serverConfig, err := GetServerConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		LogLevel: logLevel,
		Server: &ServerConfig{
			ServerPort:   serverConfig.ServerPort,
			PprofEnabled: serverConfig.PprofEnabled,
		},
	}, nil

}

func get(name string) (string, error) {
	v, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("env var %s not found", name)
	}
	return v, nil
}

func getInt(name string) (int, error) {
	v, ok := os.LookupEnv(name)
	if !ok {
		return 0, fmt.Errorf("env var %s not found", name)
	}
	intV, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("en var %s must be a number", name)
	}
	return intV, nil
}

func getBool(name string) (bool, error) {
	value, ok := os.LookupEnv(name)
	if !ok {
		return false, fmt.Errorf("env var %s not found", name)
	}
	boolV, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf("en var %s must be a bool", name)
	}
	return boolV, nil
}
