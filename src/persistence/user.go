package persistence

import (
	"context"
	"github.com/zaza-hikayat/go-rest-sample/src/app/entities"
	"github.com/zaza-hikayat/go-rest-sample/src/app/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (res entities.MemberEntity, err error) {
	err = u.db.WithContext(ctx).Where("email = ?", email).First(&res).Error

	return res, err
}

func (u userRepository) All(ctx context.Context) ([]entities.MemberEntity, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepository(db *gorm.DB) repository.IUserRepository {
	return &userRepository{db: db}
}
