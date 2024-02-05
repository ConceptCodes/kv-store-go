package config

import (
	"kv-store/pkg/logger"
	_log "log"

	"github.com/spf13/viper"
)

type Config struct {
	DefaultTTL    int    `mapstructure:"default_ttl"`
	Port          int    `mapstructure:"port"`
	Timeout       int    `mapstructure:"timeout"`
	RedisHost     string `mapstructure:"redis_host"`
	RedisPort     int    `mapstructure:"redis_port"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
}

var AppConfig *Config

func LoadAppConfig() {
	log := logger.GetLogger()
	log.Debug().Msg("Loading Server Configurations...")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		_log.Fatal(err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		_log.Fatal(err)
	}
}
