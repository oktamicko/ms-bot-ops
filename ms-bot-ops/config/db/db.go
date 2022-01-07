package db

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"

	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"

	"bitbucket.org/bridce/ms-bot-ops/config"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/qor/validations"
)

// DB Global DB connection
var DbCore *gorm.DB
var DbOCH *gorm.DB
var DbCerCore *gorm.DB
var DbMongo *mgo.Database
var DbCerSta *gorm.DB

func init() {
	if DbCore != nil || DbOCH != nil {
		return
	}
	var err error

	dbConfig := config.Config.DatabaseCore
	dbConfigOch := config.Config.DatabaseOCH
	dbConfigCerCore := config.Config.DatabaseCeriaCore
	dbConfigCerStaging := config.Config.DatabaseCeriaStaging

	if dbConfig.Adapter == "cockroach" {
		if dbConfig.SslMode == "disable" {
			DbCore, err = gorm.Open("postgres", fmt.Sprintf("postgresql://%v%v@%v:%v/%v?application_name=cockroach&sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		}
	} else if dbConfig.Adapter == "postgres" {
		if dbConfig.SslMode == "disable" {
			pg_con_string_core := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)
			DbCore, err = gorm.Open("postgres", pg_con_string_core)

			pg_con_string_cer_staging := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbConfigCerStaging.Host, dbConfigCerStaging.Port, dbConfigCerStaging.User, dbConfigCerStaging.Password, dbConfigCerStaging.Name)
			DbCerSta, err = gorm.Open("postgres", pg_con_string_cer_staging)

			pg_con_string_och := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbConfigOch.Host, dbConfigOch.Port, dbConfigOch.User, dbConfigOch.Password, dbConfigOch.Name)
			DbOCH, err = gorm.Open("postgres", pg_con_string_och)

			pg_con_string_cer_core := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				dbConfigCerCore.Host, dbConfigCerCore.Port, dbConfigCerCore.User, dbConfigCerCore.Password, dbConfigCerCore.Name)
			DbCerCore, err = gorm.Open("postgres", pg_con_string_cer_core)
		} else {
			panic(fmt.Errorf("Unsupported SSL Mode: ", dbConfig.SslMode))
		}
	}

	//MongoDB
	info := &mgo.DialInfo{
		Addrs:    []string{config.Config.DatabaseCeriaOch.Host1, config.Config.DatabaseCeriaOch.Host2, config.Config.DatabaseCeriaOch.Host3},
		Timeout:  1200 * time.Second,
		Database: config.Config.DatabaseCeriaOch.Name,
		Username: config.Config.DatabaseCeriaOch.User,
		Password: config.Config.DatabaseCeriaOch.Password,
	}

	SessionMongoDb, errM := mgo.DialWithInfo(info)
	if errM != nil {
		fmt.Println(errM)
	} else {
		fmt.Println("Session created")
	}
	DbMongo = SessionMongoDb.DB(config.Config.DatabaseCeriaOch.Name)

	if err == nil {
		DbCore.LogMode(true)
		DbOCH.LogMode(true)
		DbCerCore.LogMode(true)
		DbCerSta.LogMode(true)
		//if os.Getenv("DEBUG") != "" {
		//	DB.LogMode(true)
		//}

		validations.RegisterCallbacks(DbOCH)
		validations.RegisterCallbacks(DbCore)
		validations.RegisterCallbacks(DbCerCore)
		validations.RegisterCallbacks(DbCerSta)
	} else {
		panic(err)
	}
}
