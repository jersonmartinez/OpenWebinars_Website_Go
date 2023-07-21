package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./web/")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("No se pudo leer el archivo de configuración: %s", err)
	}

	cfg := &Config{
		Port: viper.GetString("port"),
	}

	return cfg
}
