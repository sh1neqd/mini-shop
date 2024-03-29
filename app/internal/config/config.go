package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	Listen        struct {
		Type       string `env:"LISTEN_TYPE" env-default:"port" env-description:"'port' or 'sock'. if 'sock' then env 'SOCKET_FILE' is required"`
		BindIP     string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port       string `env:"PORT" env-default:"5000"`
		SocketFile string `env:"SOCKET_FILE" env-default:"app.sock"`
	}
	AppConfig struct {
		LogLevel  string `env:"LOG_LEVEL" env-default:"trace"`
		AdminUser struct {
			Email    string `env:"ADMIN_EMAIL" env-default:"admin"`
			Password string `env:"ADMIN_PWD" env-default:"admin"`
		}
	}
	PostgreSQL struct {
		Host     string `env:"PSQL_HOST"  env-default:"db"`
		Port     int    `env:"PSQL_PORT"  env-default:"5432"`
		Username string `env:"PSQL_USERNAME"  env-default:"postgres"`
		Password string `env:"PSQL_PASSWORD" env-default:"postgres"`
		Database string `env:"PSQL_DATABASE"  env-default:"postgres"`
		SSLMode  string `env:"SSL_MODE" env-default:"disable"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("gather config")

		instance = &Config{}

		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "sh1neqd - mini-Shop"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
