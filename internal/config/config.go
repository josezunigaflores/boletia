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
	// Schedule defines the schedule that the service uses for update the sources.
	Schedule int `default:"30" required:"true" env:"SCHEDULE"`
	// TimeOut is the time that requests lives.
	TimeOut int `default:"1600" required:"true" env:"TIMEOUT"`
	//https://api.currencyapi.com/v3/latest?apikey=hJko5diT3ZVpLt5vsQ1yU6acNSipVsbr9HujSyjA
	PathCurrency string `default:"any" required:"true" env:"PATH_CURRENCY"`
}{}

func Init() error {
	if err := configor.Load(&Config, "config.json"); err != nil {
		return err
	}

	return nil
}
