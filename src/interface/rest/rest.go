package rest

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	_recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
	"github.com/zaza-hikayat/go-rest-sample/src/app/usecase"
	"github.com/zaza-hikayat/go-rest-sample/src/config"
	"github.com/zaza-hikayat/go-rest-sample/src/interface/rest/handlers"
	handler_admin "github.com/zaza-hikayat/go-rest-sample/src/interface/rest/handlers/admin"
	"github.com/zaza-hikayat/go-rest-sample/src/interface/rest/middleware"
	"github.com/zaza-hikayat/go-rest-sample/src/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type HttpServer struct {
	App              *fiber.App
	AllUsecase       usecase.AllUsecase
	GracefulShutdown func(ct context.Context) error
	JwtHelper        *utils.JWTHelper
}

func NewHttpServer(allUsecase usecase.AllUsecase, jwtHelper *utils.JWTHelper) *HttpServer {
	return &HttpServer{
		App: fiber.New(fiber.Config{
			DisableStartupMessage: false,
			ErrorHandler:          middleware.NewErrorHandler(),
		}),
		AllUsecase: allUsecase,
		JwtHelper:  jwtHelper,
	}
}

func (httpServer HttpServer) Run(conf config.Config) {
	UC := httpServer.AllUsecase
	app := httpServer.App
	app.Use(_recover.New())
	app.Use(middleware.NewLogger())
	middlewareJwt := middleware.JwtToken(httpServer.JwtHelper)

	routerV1 := app.Group("/api/v1")
	routerAdminV1 := routerV1.Group("/admin")

	authHandler := handlers.NewAuthHandler(UC.AuthUsecase)
	authHandler.Setup(routerV1)

	adminHandler := handler_admin.NewAdminHandler(UC.Admin)
	adminHandler.Setup(routerAdminV1, middleware.BasicAuth(conf), middlewareJwt)

	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", conf.Http.Port)); err != nil {
			logrus.Info("Shutting down server: %v", err)
		}
	}()

	logrus.Infof("Server is running on :%s", conf.Http.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logrus.Infof("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if httpServer.GracefulShutdown != nil {
		_ = httpServer.GracefulShutdown(ctx)
	}

	if err := app.Shutdown(); err != nil {
		logrus.Fatalf("Server shutdown failed: %v", err)
	}

	logrus.Infof("Server exited properly")
}
