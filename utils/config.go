package utils

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"gorm.io/gorm/logger"
)

type Config struct {

	// Gin mód
	GinMode string `envconfig:"GIN_MODE" required:"true"`

	// URL aplikace
	AppUrl string `envconfig:"APP_URL" required:"true"`

	// Swagger
	Swagger bool `envconfig:"SWAGGER" required:"true"`

	// Admin
	AdminMail     string `envconfig:"ADMIN_MAIL" required:"true"`
	AdminPassword string `envconfig:"ADMIN_PW" required:"true"`

	// PostgreSQL
	PostgresUser     string `envconfig:"PG_USER" required:"true"`
	PostgresPassword string `envconfig:"PG_PW" required:"true"`
	PostgresServer   string `envconfig:"PG_HOST" required:"true"`
	PostgresPort     int    `envconfig:"PG_PORT" required:"true"`
	PostgresDb       string `envconfig:"PG_DB" required:"true"`

	// Access token
	AccessTokenLifespan int    `envconfig:"ACCESS_TOKEN_LIFESPAN" required:"true"`
	AccessTokenSecret   string `envconfig:"ACCESS_TOKEN_SECRET" required:"true"`
}

// Vrátí config.
//
//	@return *Config
//	@return error
func LoadConfig() (*Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}
	return &config, nil
}

// Returns the Gorm log level according to the environment variable
//
//	@receiver c
//	@return logger.LogLevel
func (c *Config) GetGormLogLevel() logger.LogLevel {
	if c.GinMode == "release" {
		return logger.Error
	}
	return logger.Info
}
