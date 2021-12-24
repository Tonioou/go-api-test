package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Postgres struct {
		Url string
	}
	LogLevel string
}

var configRunOnce sync.Once
var config *Config

func GetConfig() *Config {
	configRunOnce.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()

		viper.SetDefault("log.level", "info")
		viper.SetDefault("postgres.host", "postgres")
		viper.SetDefault("postgres.port", 5432)
		viper.SetDefault("postgres.username", "postgres")
		viper.SetDefault("postgres.password", "postgres")
		viper.SetDefault("postgres.database", "todo_list")

		databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			viper.GetString("postgres.username"),
			viper.GetString("postgres.password"),
			viper.GetString("postgres.host"),
			viper.GetString("postgres.port"),
			viper.GetString("postgres.database"),
		)

		config = &Config{
			Postgres: struct{ Url string }{
				Url: databaseUrl,
			},
			LogLevel: viper.GetString("log.level"),
		}
	})

	return config
}
