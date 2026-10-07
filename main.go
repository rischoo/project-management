package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rischoo/project-management/config"
	"github.com/rischoo/project-management/controllers"
	"github.com/rischoo/project-management/database/seed"
	"github.com/rischoo/project-management/repositories"
	"github.com/rischoo/project-management/routes"
	"github.com/rischoo/project-management/services"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	seed.SeedAdmin()
	app := fiber.New()

	userRepo := repositories.NewuserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	routes.Setup(app, userController)

	port := config.AppConfig.AppPort
	log.Println("Server running on port : ", port)
	log.Fatal(app.Listen(":" + port))
}
