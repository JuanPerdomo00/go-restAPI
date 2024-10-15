package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DSN = "host=localhost user=admin password=YWRtaW5pc3RyYXRvcgo= dbname=dbadmin port=5432"

var DB *gorm.DB

func DBconnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB is runing")
	}
}
