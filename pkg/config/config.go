package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DefaultTTL int    `mapstructure:"defualt_ttl"`
	Port       int    `mapstructure:"port"`
	KeyRegex   string `mapstructure:"key_regex"`
	ValueRegex string `mapstructure:"value_regex"`
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
