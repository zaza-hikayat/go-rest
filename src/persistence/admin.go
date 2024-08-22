package persistence

import (
	"context"
	"github.com/zaza-hikayat/go-rest-sample/src/app/entities"
	"github.com/zaza-hikayat/go-rest-sample/src/app/repository"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) repository.IAdminRepository {
	return &adminRepository{
		db: db,
	}
}

func (r adminRepository) GetByEmail(ctx context.Context, email string) (e *entities.AdminEntity, err error) {
	err = r.db.WithContext(ctx).Where("email = ?", email).First(&e).Error

	return e, err
}

func (r adminRepository) Create(ctx context.Context, e entities.AdminEntity) (entity *entities.AdminEntity, err error) {
	err = r.db.WithContext(ctx).Create(&e).Error
	if err != nil {
		return nil, err
	}

	return &e, err
}

func (r adminRepository) Delete(ctx context.Context, id uint) (entity *entities.AdminEntity, err error) {
	err = r.db.
		Where("id = ?", id).
		Delete(&entity).
		Error

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (r adminRepository) Count(ctx context.Context) (count uint, err error) {
	var _count int64
	err = r.db.Model(&entities.AdminEntity{}).Count(&_count).Error

	return uint(_count), err
}

func (r adminRepository) List(ctx context.Context) (e []entities.AdminEntity, err error) {
	err = r.db.WithContext(ctx).Find(&e).Error

	return e, err
}
