package database

import (
	"casic/src/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *sql.DB
var DBG *gorm.DB

/*
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "987123"
	dbname   = "postgres"
)*/
const (
	host     = "192.168.0.57"
	port     = 5432 // Default port
	user     = "habrpguser"
	password = "pgpwd4habr"
	dbname   = "habrdb"
)

func Connect() error {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}

func AutoMigrate() error {
	dsn := "host=localhost user=postgres password=987123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DBG, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DBG.AutoMigrate(models.Order{}, models.OrderItem{})
	return nil
}
