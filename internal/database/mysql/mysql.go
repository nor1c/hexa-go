package mysql

import (
	"fmt"
	"gc-hexa-go/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

type MySQLConfig struct {
	Driver   string
	Username string
	Password string
	Database string
}

func Connect() {
	dbConfig := &MySQLConfig{
		Driver:   config.GetEnv("DB_DRIVER"),
		Username: config.GetEnv("DB_USER"),
		Password: config.GetEnv("DB_PASS"),
		Database: config.GetEnv("DB_NAME"),
	}

	dataSourceName := fmt.Sprintf("%s:%s@/%s", dbConfig.Username, dbConfig.Password, dbConfig.Database)

	db, err := sqlx.Connect(dbConfig.Driver, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	DB = db
}
