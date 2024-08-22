package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/zaza-hikayat/go-rest-sample/src/config"
)

func BasicAuth(conf config.Config) fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			conf.BasicAuth.Username: conf.BasicAuth.Password,
		},
		Unauthorized: func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusUnauthorized).JSON(
				fiber.Map{
					"success": false,
					"message": "Unauthorized User",
				})
		},
	})
}
