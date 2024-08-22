package utils

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/zaza-hikayat/go-rest-sample/src/app/constants"
	app_errors "github.com/zaza-hikayat/go-rest-sample/src/app/errors"
	"github.com/zaza-hikayat/go-rest-sample/src/app/repository"
	"github.com/zaza-hikayat/go-rest-sample/src/config"
)

type JWTClaims struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWTHelper struct {
	tokenSecret        string
	refreshTokenSecret string
	adminRepo          repository.IAdminRepository
}

func NewJWTHelper(config config.Config, adminRepo repository.IAdminRepository) *JWTHelper {
	return &JWTHelper{
		tokenSecret:        config.JWT.TokenSecret,
		refreshTokenSecret: config.JWT.RefreshSecret,
		adminRepo:          adminRepo,
	}
}

func (helper *JWTHelper) GenerateToken(claims JWTClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString(helper.tokenSecret)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (helper *JWTHelper) ParseToken(tokenString string) (claim *JWTClaims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return []byte(helper.tokenSecret), nil
	})

	// Checking token validity
	if !token.Valid {
		return nil, app_errors.NewError(err, "invalid token")
	}

	if claim != nil {
		switch true {
		case claim.Role == constants.ROLE_SUPERADMIN || claim.Role == constants.ROLE_ADMIN:
			ctx := context.Background()
			_, err := helper.adminRepo.GetByEmail(ctx, claim.Email)
			if err != nil {
				return nil, err
			}
			break
		default:

			break
		}
	}

	return claim, err
}
