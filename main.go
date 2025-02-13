package main

import (
	"backend_services/controllers"
	"backend_services/database"
	"backend_services/handlers/http"
	"backend_services/util"
	"fmt"
	"os"

	_ "backend_services/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/joho/godotenv"
	// "gorm.io/gorm/clause"
)

// @title FullStack Developer __(Client : KFC)
// @version 1.0
// @description FullStack Developer __(Client : KFC) Rest API.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email lifelinejar@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	} else {
		fmt.Println("JWT_SECRET:", os.Getenv("JWT_SECRET"))
	}

	database.Connect()

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*util.Error); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*fiber.Error); ok {
				return ctx.Status(e.Code).JSON(util.Error{Status: e.Code, Code: "internal-server", Message: e.Message})
			} else {
				return ctx.Status(500).JSON(util.Error{Status: 500, Code: "internal-server", Message: err.Error()})
			}
		},
	})

	// setup middleware
	app.Use(cors.New(cors.Config{
		AllowCredentials: false,
	}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Initialize controllers
	authController := controllers.NewAuthController(database.DB)
	userController := controllers.NewUserController(database.DB)

	// Register routes
	http.RegisterRoutes(app,
		authController,
		userController)

	app.Listen(":" + os.Getenv("PORT"))
}
