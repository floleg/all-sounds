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

func LoadConfig(env string) (config Config, err error) {
	viper.AddConfigPath("./configs")
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
