package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"log"
)

type appConfig struct {
	AppPort          string `env:"APP_PORT"`
	EntryHost        string `env:"ENTRY_HOST"`
	SourceDbHost     string `env:"SOURCE_DB_HOST"`
	SourceDbPort     string `env:"SOURCE_DB_PORT"`
	SourceDbUser     string `env:"SOURCE_DB_USER"`
	SourceDbPassword string `env:"SOURCE_DB_PASSWORD"`
	SourceDbDbname   string `env:"SOURCE_DB_DBNAME"`
	SourceDbDriver   string `env:"SOURCE_DB_DRIVER"`
}

var AppCfg appConfig

func Init() {

	// ENV config for local/development purposes
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	// Default config suitable for Docker side configuration
	appCfg := appConfig{
		AppPort:          "8889",
		EntryHost:        "0.0.0.0",
		SourceDbHost:     "PgSQL-DB-Source",
		SourceDbPort:     "5432",
		SourceDbUser:     "docker_test",
		SourceDbPassword: "docker_test",
		SourceDbDbname:   "docker_test",
		SourceDbDriver:   "postgres",
	}

	if errParse := env.Parse(&appCfg); errParse != nil {
		log.Println(errParse)
	}

	AppCfg = appCfg
}
