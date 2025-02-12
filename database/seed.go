package database

import (
	"log"

	"gorm.io/gorm"
)

func RunSeeders(db *gorm.DB) {
	// seeder.SeedUsers(db)
	log.Println("Seeding completed.")
}
