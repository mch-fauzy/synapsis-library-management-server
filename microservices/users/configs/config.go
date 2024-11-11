package configs

import (
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Config is a struct that will receive configuration options via environment variables.
type Config struct {
	App struct {
		JwtAccessKey string `mapstructure:"JWT_ACCESS_KEY"`
	}

	DB struct {
		PostgreSQL struct {
			Host     string `mapstructure:"HOST"`
			Port     string `mapstructure:"PORT"`
			Username string `mapstructure:"USER"`
			Password string `mapstructure:"PASSWORD"`
			Name     string `mapstructure:"NAME"`
			Timezone string `mapstructure:"TIMEZONE"`
		}
	}

	Server struct {
		Port string `mapstructure:"PORT"`
	}
}

var (
	conf Config
	once sync.Once
)

// Get are responsible to load env and get data an return the struct
func Get() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal().Err(err).Msg("Failed reading config file")
	}

	once.Do(func() {
		log.Info().Msg("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
	})

	return &conf
}
