package middleware

import (
	"Search/util/loaddotenv"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"strings"
)

func FiberMiddleware(app *fiber.App) {
	allowOrigins := strings.Join([]string{
		loaddotenv.LoadDotEnv("LOCAL_SERVER"),
		loaddotenv.LoadDotEnv("EC2_SERVER"),
	}, ", ")

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: fiber.MethodGet,
		AllowHeaders: "*",
	}))
}
