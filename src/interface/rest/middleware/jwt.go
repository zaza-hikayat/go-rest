package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zaza-hikayat/go-rest-sample/src/utils"
)

func JwtToken(jwtHelper *utils.JWTHelper) fiber.Handler {
	return func(c *fiber.Ctx) error {

	}
}
