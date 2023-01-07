package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host string `mapstructure:"HOST"`
}

type DbConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	Database string `mapstructure:"DB_NAME"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
}

type FrontConfig struct {
	Host string `mapstructure:"FRONT_HOST"`
}

type Config struct {
	ServerConfig ServerConfig
	DbConfig     DbConfig
	FrontConfig  FrontConfig
}

func SetupSource() {
	viper.SetConfigType("env")
	viper.AutomaticEnv()
}

func NewConfig() *Config {
	SetupSource()

	return &Config{
		ServerConfig: ServerConfig{viper.GetString("HOST")},
		DbConfig: DbConfig{
			viper.GetString("DB_HOST"),
			viper.GetString("DB_NAME"),
			viper.GetString("DB_USER"),
			viper.GetString("DB_PASSWORD"),
		},
		FrontConfig: FrontConfig{
			viper.GetString("FRONT_HOST"),
		},
	}
}

func Logger() (logger zerolog.Logger) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	return zerolog.New(output).With().Timestamp().Logger()
}
