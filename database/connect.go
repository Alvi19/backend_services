package database

import (
	"backend_services/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "public.", // this is the schema
			SingularTable: true,
		},
	})

	if err != nil {
		panic(fmt.Errorf("Fatal error connect DB: %w \n", err))
	}

	fmt.Println("Db connect success")
	DB = db

	db.AutoMigrate(&models.User{})
}
