package config

import "github.com/jinzhu/configor"

var Config = struct {
	Port            uint   `default:"8080" required:"true" env:"PORT"`
	Host            string `default:"localhost" required:"true" env:"HOST"`
	DBUser          string `default:"developer" required:"true" env:"DB_USER"`
	DBPass          string `default:"Clever.Dev" required:"true" env:"DB_PASS"`
	DBHost          string `default:"localhost" required:"true" env:"DB_HOST"`
	DBPort          string `default:"5432" required:"true" env:"DB_PORT"`
	DBName          string `default:"clever" required:"true" env:"DB_NAME"`
	GinMode         string `default:"debug" required:"true" env:"GIN_MODE"`
	ShutdownTimeout int    `default:"30" required:"true" env:"SHUTDOWN_TIMEOUT"`
}{}

func Init() error {
	if err := configor.Load(&Config); err != nil {
		return err
	}

	return nil
}
