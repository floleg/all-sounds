// Package config parses the yaml configuration files
// based on the runtime defined ENV variable.
package config

import (
	"fmt"

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

// LoadConfig takes the ENV variable as an arguments to load
// the proper configuration yaml file and map its values to
// config.Config struct instance
func LoadConfig(env string, configPath string) (config Config, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("LoadConfig: %w", err))
	}

	err = viper.Unmarshal(&config)
	return
}
