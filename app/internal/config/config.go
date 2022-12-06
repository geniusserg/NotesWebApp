package config

import (
	"cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"port"`
	Listen        struct {
		Type   string `env:"LISTEN_TYPE" env-default:"port"`
		BindIP string `env:"BIND_IP" env-default:"0.0.0.0"`
		Port   string `env:"PORT" env-default:"10000"`
	}
	AppConfig struct {
		LogLevel  string
		AdminUser struct {
			Email    string
			Password string
		}
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Println("Gather config")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			help := "System error - environment vars broken"
			help, _ = cleanenv.GetDescription(instance, &help)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
