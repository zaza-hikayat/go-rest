package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Setup(router fiber.Router, middleware ...fiber.Handler)
}
