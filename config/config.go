package config

import (
	"gosampleapi/singletons"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port   uint16 `mapstructure:"PORT"`
	Secret string `mapstructure:"SECRET" validate:"required,min=32"`

	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

func GetConfig() *Config {
	var config Config

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return nil
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println(err)
		return nil
	}

	err = singletons.Validate.Struct(config)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &config
}
