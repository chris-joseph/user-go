package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/pkg/errors"
	"github.com/samber/lo"
)

// constants
const (
	DevelopmentEnvironment = "Development"
	StagingEnvironment     = "Staging"
	ProductionEnvironment  = "Production"
)

type EnvironmentVariables struct {
	Environment     string `env:"ENV,required"`
	Port            int64  `env:"PORT,required"`
	Debug           bool   `env:"DEBUG,required"`
	ApplicationName string `env:"APPLICATION_NAME,required"`

	MongoHostURI string `env:"MONGO_HOST_URI,required"`
	MongoDBName  string `env:"MONGO_DB_NAME,required"`

	RedisHost string `env:"REDIS_HOST,required"`
	RedisPort string `env:"REDIS_PORT,required"`
}

type MongoConfig struct {
	HostURI string
	DBName  string
}

type RedisConfig struct {
	Host string
	Port string
}

type Config struct {
	Environment     string
	Port            int64
	Debug           bool
	ApplicationName string
	Mongo           MongoConfig
	Redis           RedisConfig
}

func extractEnvironmentVariables() (*EnvironmentVariables, error) {
	envVariables := EnvironmentVariables{}
	if err := env.Parse(&envVariables); err != nil {
		return nil, err // some required environment variables are missing
	}
	return &envVariables, nil
}

func NewConfig() (*Config, error) {
	env, err := extractEnvironmentVariables()
	if err != nil {
		return nil, err
	}
	if !lo.Contains([]string{DevelopmentEnvironment, StagingEnvironment, ProductionEnvironment}, env.Environment) {
		return nil, errors.New(fmt.Sprintf("Invalid value '%s' passed for `ENVIRONMENT`", env.Environment))
	}
	return &Config{
		Environment:     env.Environment,
		Port:            env.Port,
		Debug:           env.Debug,
		ApplicationName: env.ApplicationName,
		Mongo: MongoConfig{
			HostURI: env.MongoHostURI,
			DBName:  env.MongoDBName,
		},
		Redis: RedisConfig{
			Host: env.RedisHost,
			Port: env.RedisPort,
		},
	}, nil
}
