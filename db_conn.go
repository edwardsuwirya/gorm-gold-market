package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DbConn struct {
	Db *gorm.DB
}

func (d *DbConn) PingTest() {
	db, err := d.Db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Database connected")
	}
}

func (d *DbConn) Migration(tables ...interface{}) {
	err := d.Db.AutoMigrate(tables...)
	if err != nil {
		log.Fatalln(err)
	}
}
func (d *DbConn) Close() {
	db, err := d.Db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func NewDbConn() *DbConn {
	dbHost := "159.223.42.164"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "P@ssw0rd"
	dbName := "gold_market"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	env := "dev"
	//Daftar konfigurasi GORM apa saja yang ada
	//https://gorm.io/docs/gorm_config.html
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if env == "dev" {
		return &DbConn{
			Db: db.Debug(),
		}
	}

	return &DbConn{
		Db: db,
	}

}
