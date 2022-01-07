package config

import (
	"os"

	"github.com/caarlos0/env"
	//_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type DatabaseCore struct {
	Name          string `env:"DB_SCHEMA" default:"direct_debit"`
	Adapter       string `env:"DB_DRIVER" default:"mysql"`
	Host          string `env:"DB_HOST" default:"localhost"`
	Port          string `env:"DB_PORT" default:"3306"`
	User          string `env:"DB_USER"`
	Password      string `env:"DB_PASSWORD"`
	SslMode       string `env:"DB_SSL_MODE"`
	TableCharges  string `env:"TABLE_CHARGES"`
	TableRefunds  string `env:"TABLE_REFUNDS"`
	SchedulerTime string `env:"SCH_TIME"`
	SchedulerMode string `env:"SCH_MODE"`
}

type DatabaseStaging struct {
	Name          string `env:"DB_SCHEMA_STAGING" default:"direct_debit"`
	Adapter       string `env:"DB_DRIVER_STAGING" default:"mysql"`
	Host          string `env:"DB_HOST_STAGING" default:"localhost"`
	Port          string `env:"DB_PORT_STAGING" default:"3306"`
	User          string `env:"DB_USER_STAGING"`
	Password      string `env:"DB_PASSWORD_STAGING"`
	SslMode       string `env:"DB_SSL_MODE_STAGING"`
	TableCharges  string `env:"TABLE_CHARGES_STAGING"`
	TableRefunds  string `env:"TABLE_REFUNDS_STAGING"`
	SchedulerTime string `env:"SCH_TIME_STAGING"`
	SchedulerMode string `env:"SCH_MODE_STAGING"`
}

type DatabaseOCH struct {
	Name     string `env:"DB_SCHEMA_OCH" default:"direct_debit"`
	Adapter  string `env:"DB_DRIVER_OCH" default:"mysql"`
	Host     string `env:"DB_HOST_OCH" default:"localhost"`
	Port     string `env:"DB_PORT_OCH" default:"3306"`
	User     string `env:"DB_USER_OCH"`
	Password string `env:"DB_PASSWORD_OCH"`
	SslMode  string `env:"DB_SSL_MODE_OCH"`
}

type DatabaseCeriaCore struct {
	Name     string `env:"DB_SCHEMA_CERIA_CORE" default:"direct_debit"`
	Adapter  string `env:"DB_DRIVER_CERIA_CORE" default:"mysql"`
	Host     string `env:"DB_HOST_CERIA_CORE" default:"localhost"`
	Port     string `env:"DB_PORT_CERIA_CORE" default:"3306"`
	User     string `env:"DB_USER_CERIA_CORE"`
	Password string `env:"DB_PASSWORD_CERIA_CORE"`
	SslMode  string `env:"DB_SSL_MODE_CERIA_CORE"`
}

type DatabaseCeriaOch struct {
	Name    string `env:"DB_SCHEMA_CERIA_OCH" default:"direct_debit"`
	Adapter string `env:"DB_DRIVER_CERIA_OCH" default:"mysql"`
	Host1   string `env:"DB_HOST_CERIA_OCH1" default:"localhost"`
	// Host2    string `env:"DB_HOST_CERIA_OCH2" default:"localhost"`
	// Host3    string `env:"DB_HOST_CERIA_OCH3" default:"localhost"`
	Port     string `env:"DB_PORT_CERIA_OCH" default:"3306"`
	User     string `env:"DB_USER_CERIA_OCH"`
	Password string `env:"DB_PASSWORD_CERIA_OCH"`
	SslMode  string `env:"DB_SSL_MODE_CERIA_OCH"`
}

type ServerConfig struct {
	ServiceName          string `env:"SERVICE_NAME,required"`
	ServiceVersion       string `env:"SERVICE_VERSION,required"`
	ServicePort          string `env:"SERVICE_PORT,required" envDefault:"7777"`
	ServiceHost          string `env:"SERVICE_HOST,required"`
	HTTPProxy            string `env:"HTTP_PROXY"`
	DatabaseCore         DatabaseCore
	DatabaseOCH          DatabaseOCH
	DatabaseCeriaCore    DatabaseCeriaCore
	DatabaseCeriaOch     DatabaseCeriaOch
	DatabaseCeriaStaging DatabaseStaging
}

var Config ServerConfig

func init() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}
}

func loadConfig() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Warn().Msg("Cannot find .env file. OS Environments will be used")
	}
	err = env.Parse(&Config)
	err = env.Parse(&Config.DatabaseOCH)
	err = env.Parse(&Config.DatabaseCore)
	err = env.Parse(&Config.DatabaseCeriaCore)
	err = env.Parse(&Config.DatabaseCeriaOch)
	err = env.Parse(&Config.DatabaseCeriaStaging)
	return err
}

func IsDevelopmentMode() bool {
	return os.Getenv("ENV") == "DEVELOPMENT"
}

func IsDebugMode() bool {
	return os.Getenv("ENV") == "DEBUG" || IsDevelopmentMode()
}

func IsVerifyActive() bool {
	return os.Getenv("DD_VERIFY_SIGNATURE") == "1"
}
