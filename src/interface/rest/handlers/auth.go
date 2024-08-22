package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zaza-hikayat/go-rest-sample/src/app/dto"
	app_errors "github.com/zaza-hikayat/go-rest-sample/src/app/errors"
	"github.com/zaza-hikayat/go-rest-sample/src/app/usecase"
)

type IAuthHandler interface {
	Handler
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	ForgotPassword(c *fiber.Ctx) error
}

type authHandler struct {
	authUsecase usecase.IAuthUsecase
}

func (a authHandler) Setup(router fiber.Router, middleware ...fiber.Handler) {
	router.Post("login", a.Login)
	router.Post("logout", a.Logout)
	router.Post("forgot-password", a.ForgotPassword)
}

func (a authHandler) ForgotPassword(c *fiber.Ctx) error {
	req := new(dto.LoginReqDTO)

	if err := c.BodyParser(req); err != nil {
		return app_errors.NewError(err, "failed to parse request body")
	}

	if err := req.Validate(); err != nil {
		return err
	}

	result, err := a.authUsecase.Login(c, *req)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (a authHandler) Login(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a authHandler) Logout(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthHandler(authUsecase usecase.IAuthUsecase) IAuthHandler {
	return &authHandler{authUsecase: authUsecase}
}
