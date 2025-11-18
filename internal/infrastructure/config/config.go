package config

import (
	"os"
	"reflect"

	"github.com/arfanxn/welding/pkg/reflectutil"
	"github.com/joho/godotenv"
)

type Config struct {
	// App
	AppName string `env:"APP_NAME"`
	AppHost string `env:"APP_HOST"`
	AppPort string `env:"APP_PORT"`

	// Gin
	GinMode string `env:"GIN_MODE"`

	// Database
	PostgresDSN string `env:"POSTGRES_DSN"`

	// Log
	LogLevel    string `env:"LOG_LEVEL"`
	LogFilepath string `env:"LOG_FILEPATH"`

	// JWT
	JWTSecret   string `env:"JWT_SECRET"`
	JWTDuration int    `env:"JWT_DURATION"`

	// Mail
	MailHost        string `env:"MAIL_HOST"`
	MailPort        int    `env:"MAIL_PORT"`
	MailIdentity    string `env:"MAIL_IDENTITY"`
	MailUsername    string `env:"MAIL_USERNAME"`
	MailPassword    string `env:"MAIL_PASSWORD"`
	MailEncryption  string `env:"MAIL_ENCRYPTION"`
	MailFromAddress string `env:"MAIL_FROM_ADDRESS"`
	MailFromName    string `env:"MAIL_FROM_NAME"`
}

// NewConfigFromEnv creates a new Config instance with values from environment variables
func NewConfigFromEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config := &Config{}

	val := reflect.ValueOf(config).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		envKey := field.Tag.Get("env")
		if envKey == "" {
			continue
		}

		envValue, exists := os.LookupEnv(envKey)
		if !exists {
			continue
		}

		if err := reflectutil.SetValueFromString(val.Field(i), envValue); err != nil {
			return nil, err
		}
	}

	return config, nil
}
