package database

import "fmt"

// DatabaseConfig configurations of database connection
type DatabaseConfig struct {
	Host            string `mapstructure:"DB_HOST"`
	Port            string `mapstructure:"DB_PORT"`
	UserName        string `mapstructure:"DB_USERNAME"`
	Password        string `mapstructure:"DB_PASSWORD"`
	Name            string `mapstructure:"DB_NAME"`
	ConnectAttempts int    `mapstructure:"DB_CONNECT_ATTEMPS"`
}

func (config DatabaseConfig) ToConnectionURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.UserName,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
}
