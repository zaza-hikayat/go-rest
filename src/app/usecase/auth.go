package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zaza-hikayat/go-rest-sample/src/app/dto"
	app_errors "github.com/zaza-hikayat/go-rest-sample/src/app/errors"
	"github.com/zaza-hikayat/go-rest-sample/src/app/repository"
)

type IAuthUsecase interface {
	Login(*fiber.Ctx, dto.LoginReqDTO) (dto.LoginResDTO, error)
	Logout(*fiber.Ctx) error
	Register(*fiber.Ctx, dto.RegisterReqDTO) error
	ForgotPassword(*fiber.Ctx, dto.ForgotPasswordReqDTO) error
}

type authUsecase struct {
	UserRepository repository.IUserRepository
}

func (a authUsecase) Login(ctx *fiber.Ctx, reqDTO dto.LoginReqDTO) (res dto.LoginResDTO, err error) {
	c := ctx.Context()
	user, err := a.UserRepository.FindByEmail(c, reqDTO.Email)
	if err != nil {
		err := app_errors.NewError(err, "failed get user")

		return res, err
	}
	_ = user
	return res, err
}

func (a authUsecase) Logout(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a authUsecase) Register(ctx *fiber.Ctx, reqDTO dto.RegisterReqDTO) error {
	//TODO implement me
	panic("implement me")
}

func (a authUsecase) ForgotPassword(ctx *fiber.Ctx, reqDTO dto.ForgotPasswordReqDTO) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthUsecase(userRepository repository.IUserRepository) IAuthUsecase {
	return &authUsecase{
		UserRepository: userRepository,
	}
}
