package http

import (
	"backend_services/controllers"
	"backend_services/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App,
	authController *controllers.AuthController,
	userController *controllers.UserControllers) {

	authGroup := app.Group("/auth")
	authGroup.Post("/login", authController.Login)
	authGroup.Post("/register", authController.Register)
	authGroup.Get("/profile", middleware.AuthMiddleware(userController.DB), authController.GetProfile)
	authGroup.Post("/logout", authController.Logout)
}
