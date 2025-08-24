package config

import (
	"log"
	"os"
	"strconv"
)

// Environment represents the application environment
type Environment struct {
	Mode           string
	Port           string
	DynamoDBLocal  bool
	LogLevel       string
	AllowedOrigins []string
}

// LoadEnvironment loads configuration from environment variables
func LoadEnvironment() *Environment {
	env := &Environment{
		Mode:           getEnv("ENV", "development"),
		Port:           getEnv("PORT", "8080"),
		DynamoDBLocal:  getBoolEnv("DYNAMODB_LOCAL", true),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		AllowedOrigins: []string{"http://localhost:3000", "http://localhost:8080"},
	}

	log.Printf("Environment loaded: %+v", env)
	return env
}

// IsProduction returns true if running in production mode
func (e *Environment) IsProduction() bool {
	return e.Mode == "production"
}

// IsDevelopment returns true if running in development mode
func (e *Environment) IsDevelopment() bool {
	return e.Mode == "development"
}

// Helper functions
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getBoolEnv(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		parsed, err := strconv.ParseBool(value)
		if err != nil {
			log.Printf("Invalid boolean value for %s: %s, using default: %t", key, value, defaultValue)
			return defaultValue
		}
		return parsed
	}
	return defaultValue
}
