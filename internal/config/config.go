package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Configs struct {
	Postgres struct {
		Url string
	}
	LogLevel    string
	ServiceName string
}

var configRunOnce sync.Once
var configs *Configs
var config = viper.New()

func GetConfig() *Configs {
	configRunOnce.Do(func() {
		config.SetDefault("service.name", "go-todo-list")
		config.SetDefault("log.level", "info")
		config.SetDefault("postgres.host", "postgres")
		config.SetDefault("postgres.port", 5432)
		config.SetDefault("postgres.username", "teste")
		config.SetDefault("postgres.password", "test")
		config.SetDefault("postgres.database", "todo_list")

		config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		config.AutomaticEnv()

		databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			config.GetString("postgres.username"),
			config.GetString("postgres.password"),
			config.GetString("postgres.host"),
			config.GetString("postgres.port"),
			config.GetString("postgres.database"),
		)

		configs = &Configs{
			Postgres: struct{ Url string }{
				Url: databaseUrl,
			},
			LogLevel:    viper.GetString("log.level"),
			ServiceName: viper.GetString("service.name"),
		}
	})

	return configs
}
