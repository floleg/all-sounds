// Package config parses the yaml configuration files
// based on the runtime defined ENV variable.
package config

import (
	"github.com/spf13/viper"
)

// Config structure maps the configuration keys and values declared
// in the repository yaml config files
type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUSer     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

// LoadConfig loads config from environment
func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()

	viper.SetDefault("DB_HOST", "")
	viper.SetDefault("DB_PORT", "")
	viper.SetDefault("DB_USER", "")
	viper.SetDefault("DB_PASSWORD", "")
	viper.SetDefault("DB_NAME", "")

	err = viper.Unmarshal(&config)
	return
}
