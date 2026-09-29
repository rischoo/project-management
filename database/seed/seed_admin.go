package seed

import (
	"log"

	"github.com/rischoo/project-management/config"
	"github.com/rischoo/project-management/models"
	"github.com/rischoo/project-management/utils"
)

func SeedAdmin() {
	password, _ := utils.HashPassword("admin123")

	admin := models.User{
		Name:     "Super Admin",
		Email:    "admin@example.com",
		Password: password,
		Role:     "admin",
	}
	if err := config.DB.FirstOrCreate(&admin, models.User{Email: admin.Email}).Error; err != nil {
		log.Println("Failed to seed admin user", err)
	} else {
		log.Println("Admin user seeded successfully")
	}
}
