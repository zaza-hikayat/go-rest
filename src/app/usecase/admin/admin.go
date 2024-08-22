package usecase_admin

import (
	"context"
	"errors"
	"fmt"
	"github.com/zaza-hikayat/go-rest-sample/src/app/constants"
	"github.com/zaza-hikayat/go-rest-sample/src/app/dto"
	"github.com/zaza-hikayat/go-rest-sample/src/app/entities"
	app_errors "github.com/zaza-hikayat/go-rest-sample/src/app/errors"
	"github.com/zaza-hikayat/go-rest-sample/src/app/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type IAdminUsecase interface {
	RegisterSuperAdmin(ctx context.Context, dto dto.RegisterSuperAdminDTO) error
	RegisterAdmin(ctx context.Context, dto dto.RegisterAdminDTO) error
	List(ctx context.Context) (list []entities.AdminEntity, err error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, dto dto.RegisterAdminDTO) error
	Login(ctx context.Context, dto dto.LoginAdminReqDTO) (dto.LoginAdminRestDTO, error)
	ValidateAdmin(ctx context.Context, email string) (bool, error)
}

var prefixXID = "IADMIN"

type adminUsecase struct {
	adminRepo repository.IAdminRepository
}

func (u adminUsecase) ValidateAdmin(ctx context.Context, email string) (res bool, err error) {
	res = true
	_, err = u.adminRepo.GetByEmail(ctx, email)
	if err != nil {
		return false, err
	}

	return res, nil
}

func NewAdminUsecase(adminRepo repository.IAdminRepository) IAdminUsecase {
	return &adminUsecase{
		adminRepo: adminRepo,
	}
}

func (u adminUsecase) RegisterSuperAdmin(ctx context.Context, dto dto.RegisterSuperAdminDTO) error {
	countAdminData, err := u.adminRepo.Count(ctx)
	if err != nil {
		return err
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	e := entities.AdminEntity{
		Xid:      fmt.Sprintf("%s%v", prefixXID, countAdminData+1),
		Password: string(passwordHash),
		Email:    dto.Email,
		Phone:    dto.Phone,
		Fullname: dto.Fullname,
		Role:     constants.ROLE_SUPERADMIN,
	}
	_, err = u.adminRepo.Create(ctx, e)

	return err
}

func (u adminUsecase) RegisterAdmin(ctx context.Context, dto dto.RegisterAdminDTO) error {
	countAdminData, err := u.adminRepo.Count(ctx)
	if err != nil {
		return err
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	e := entities.AdminEntity{
		Xid:      fmt.Sprintf("%s%v", prefixXID, countAdminData+1),
		Password: string(passwordHash),
		Email:    dto.Email,
		Phone:    dto.Phone,
		Fullname: dto.Fullname,
		Role:     constants.ROLE_ADMIN,
	}
	_, err = u.adminRepo.Create(ctx, e)

	return err
}

func (u adminUsecase) List(ctx context.Context) (list []entities.AdminEntity, err error) {
	list, err = u.adminRepo.List(ctx)

	return list, err
}

func (u adminUsecase) Delete(ctx context.Context, id uint) (err error) {
	_, err = u.adminRepo.Delete(ctx, id)

	return err
}

func (u adminUsecase) Update(ctx context.Context, id uint, dto dto.RegisterAdminDTO) error {
	_, err := u.adminRepo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := app_errors.NewError(err, "Admin not found")
			err.HttpCode = http.StatusNotFound
			return err
		}

		return err
	}

	return err
}

func (u adminUsecase) Login(ctx context.Context, dto dto.LoginAdminReqDTO) (res dto.LoginAdminRestDTO, err error) {
	admin, err := u.adminRepo.GetByEmail(ctx, dto.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err := app_errors.NewError(err, "Admin not found")
			err.HttpCode = http.StatusNotFound
			return res, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(dto.Password))
	if err != nil {
		err := app_errors.NewError(err, "Invalid password")
		err.HttpCode = http.StatusBadRequest
		return res, err
	}

	return
}
