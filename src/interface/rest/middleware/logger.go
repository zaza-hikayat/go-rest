package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"time"
)

func NewLogger() fiber.Handler {
	logger := logrus.New()

	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		stop := time.Now()
		latency := stop.Sub(start)

		method := c.Method()
		path := c.OriginalURL()
		statusCode := c.Response().StatusCode()

		logWithField := logger.WithFields(logrus.Fields{
			"status_code": statusCode,
			"latency":     latency,
			"method":      method,
			"path":        path,
		})

		logWithField.Info()
		return err
	}
}
