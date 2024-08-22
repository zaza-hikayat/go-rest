package repository

import (
	"context"
	"github.com/zaza-hikayat/go-rest-sample/src/app/entities"
)

type IAdminRepository interface {
	GetByEmail(ctx context.Context, email string) (e *entities.AdminEntity, err error)
	GetById(ctx context.Context, id uint) (e *entities.AdminEntity, err error)
	Create(ctx context.Context, e entities.AdminEntity) (entity *entities.AdminEntity, err error)
	Delete(ctx context.Context, id uint) (entity *entities.AdminEntity, err error)
	Count(ctx context.Context) (count uint, err error)
	List(ctx context.Context) (list []entities.AdminEntity, err error)
}
