package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DefaultTTL uint `mapstructure:"defualt_ttl"`
	Port       int  `mapstructure:"port"`
	Timeout    int  `mapstructure:"timeout"`
}

var AppConfig *Config

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatal(err)
	}
}
