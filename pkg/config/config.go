// Package config parses the yaml configuration files
// based on the runtime defined ENV variable.
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUSer     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

func LoadConfig(env string, configpath string) (config Config, err error) {
	viper.AddConfigPath(configpath)
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
