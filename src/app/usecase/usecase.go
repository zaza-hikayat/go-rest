package usecase

import usecase_admin "github.com/zaza-hikayat/go-rest-sample/src/app/usecase/admin"

type (
	AdminUsecase struct {
		AdminUsecase usecase_admin.IAdminUsecase
	}
	AllUsecase struct {
		AuthUsecase IAuthUsecase
		Admin       AdminUsecase
	}
)
