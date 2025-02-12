package seeder

// import (
// 	"log"
// 	"time"

// 	"backend_services/models"

// 	"github.com/google/uuid"
// 	"golang.org/x/crypto/bcrypt"
// 	"gorm.io/gorm"
// )

// func SeedUsers(db *gorm.DB) {
// 	tx := db.Begin()
// 	if tx.Error != nil {
// 		log.Printf("Error starting transaction: %v", tx.Error)
// 		return
// 	}

// 	var count int64
// 	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
// 		log.Printf("Error checking existing records: %v", err)
// 		tx.Rollback()
// 		return
// 	}

// 	if count > 0 {
// 		log.Println("Data already exists, skipping seeding.")
// 		tx.Commit()
// 		return
// 	}

// 	users := []models.User{
// 		{
// 			RoleID:          1,
// 			Name:            "John Doe",
// 			Bio:             "Admin of the system",
// 			ProfilePicURL:   "",
// 			Email:           "john.doe@example.com",
// 			Phone:           "+621234567890",
// 			Password:        "admin123",
// 			CreatedAt:       time.Now(),
// 			UpdatedAt:       time.Now(),
// 			CreatedByUserID: 1,
// 			UpdatedByUserID: 1,
// 			DeletedByUserID: 0,
// 		},
// 	}

// 	for _, user := range users {
// 		user.UUID = uuid.New().String()
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			log.Printf("Error hashing password: %v", err)
// 			tx.Rollback()
// 			return
// 		}
// 		user.Password = string(hashedPassword)
// 		if err := db.Create(&user).Error; err != nil {
// 			log.Printf("Error seeding user: %v", err)
// 			tx.Rollback()
// 			return
// 		}
// 	}

// 	log.Println("Users table seeded successfully.")
// }
