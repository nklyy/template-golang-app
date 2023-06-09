package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Environment string `required:"true" envconfig:"APP_ENV"`
	Port        string `required:"true" envconfig:"PORT"`

	MongoDb
}

type MongoDb struct {
	MongoDbName string `required:"true" envconfig:"MONGO_DB_NAME"`
	MongoDbUrl  string `required:"true" envconfig:"MONGO_DB_URL"`
}

var (
	once   sync.Once
	config *Config
)

func GetConfig() (*Config, error) {
	var err error
	once.Do(func() {
		var cfg Config
		// If you run it locally and through terminal please set up this in Load function (../.env)
		_ = godotenv.Load(".env")

		if err = envconfig.Process("", &cfg); err != nil {
			return
		}

		config = &cfg
	})

	return config, err
}
