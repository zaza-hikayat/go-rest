package handler_admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zaza-hikayat/go-rest-sample/src/app/dto"
	app_errors "github.com/zaza-hikayat/go-rest-sample/src/app/errors"
	"github.com/zaza-hikayat/go-rest-sample/src/app/usecase"
	"github.com/zaza-hikayat/go-rest-sample/src/interface/rest/handlers"
)

type IAdminHandler interface {
	handlers.Handler
	Login(c *fiber.Ctx) error
	RegisterSuperadmin(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type adminHandler struct {
	uc usecase.AdminUsecase
}

func (h adminHandler) Setup(router fiber.Router, middleware ...fiber.Handler) {
	middlewareBasicAuth := middleware[0]
	//middlewareJwtAuth := middleware[1]

	router.Post("/register/superadmin", middlewareBasicAuth, h.RegisterSuperadmin)
	router.Post("/register", middlewareBasicAuth, h.Register)

}

func (h adminHandler) Login(c *fiber.Ctx) error {
	req := new(dto.LoginAdminReqDTO)

	if err := c.BodyParser(req); err != nil {
		return app_errors.NewError(err, "failed to parse request body")
	}

	if err := req.Validate(); err != nil {
		return err
	}

	login, err := h.uc.AdminUsecase.Login(c.Context(), *req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data":    login,
	})
}

func (h adminHandler) RegisterSuperadmin(c *fiber.Ctx) error {
	req := new(dto.RegisterSuperAdminDTO)
	ctx := c.Context()
	if err := c.BodyParser(req); err != nil {
		return app_errors.NewError(err, "failed to parse request body")
	}

	if err := req.Validate(); err != nil {
		return err
	}

	err := h.uc.AdminUsecase.RegisterSuperAdmin(ctx, *req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Super Admin successfully registered",
	})
}

func (h adminHandler) Register(c *fiber.Ctx) error {
	req := new(dto.RegisterAdminDTO)
	ctx := c.Context()
	if err := c.BodyParser(req); err != nil {
		return app_errors.NewError(err, "failed to parse request body")
	}

	if err := req.Validate(); err != nil {
		return err
	}

	err := h.uc.AdminUsecase.RegisterAdmin(ctx, *req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Admin successfully registered",
	})
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase) IAdminHandler {
	return &adminHandler{
		uc: adminUsecase,
	}
}
