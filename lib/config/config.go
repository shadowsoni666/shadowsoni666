package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment         string        `mapstructure:"ENV"`
	LogLevel            string        `mapstructure:"LOG_LEVEL"`
	DBConnString        string        `mapstructure:"DB_CONN_STRING"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	PASETOSecret        string        `mapstructure:"PASETO_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RedisAddress        string        `mapstructure:"REDIS_CONN_STRING"`
	RedisPassword       string        `mapstructure:"REDIS_PW"`
}

// Load reads configuration from file or environment variables.
func Load(path string) (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Viper couldn't read in the config file. %v", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Viper could not unmarshal the configuration. %v", err)
	}
	return
}
