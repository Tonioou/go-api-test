package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type (
	Database struct {
		Name string `yaml:"name"`
	}
	Replica struct {
		Host     string   `yaml:"host"`
		Port     string   `yaml:"port"`
		Username string   `yaml:"username"`
		Password string   `yaml:"password"`
		Database Database `yaml:"database"`
	}
	Configs struct {
		Postgres struct {
			RW Replica `yaml:"rw"`
			RO Replica `yaml:"ro"`
		} `yaml:"postgres"`
		LogLevel string
		Service  struct {
			Name string `yaml:"name"`
			Log  struct {
				Level string `yaml:"level"`
			} `yaml:"log"`
			Env string `yaml:"env"`
		} `yaml:"service"`
		Otel struct {
			Exporter struct {
				GRPC struct {
					Endpoint string `yaml:"endpoint"`
				} `yaml:"grpc"`
			} `yaml:"exporter"`
		} `yaml:"otel"`
	}
)

func NewConfig() *Configs {
	var config *Configs

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	viper.SetConfigName("config")
	if err := viper.MergeInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Fatal(err)
		}
	}

	viper.SetConfigName("config-local")
	if err := viper.MergeInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Fatal(err)
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
