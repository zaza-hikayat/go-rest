package repository

import (
	"context"
	"github.com/zaza-hikayat/go-rest-sample/src/app/entities"
)

type IUserRepository interface {
	FindByEmail(ctx context.Context, email string) (entities.MemberEntity, error)
	All(context.Context) ([]entities.MemberEntity, error)
}
