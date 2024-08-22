package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/zaza-hikayat/go-rest-sample/src/app/usecase"
	usecase_admin "github.com/zaza-hikayat/go-rest-sample/src/app/usecase/admin"
	"github.com/zaza-hikayat/go-rest-sample/src/config"
	"github.com/zaza-hikayat/go-rest-sample/src/interface/rest"
	"github.com/zaza-hikayat/go-rest-sample/src/persistence"
	"github.com/zaza-hikayat/go-rest-sample/src/utils"
)

func main() {
	conf := config.Make()
	db := persistence.NewPostgres(conf)

	handlerShutdown := func(ct context.Context) error {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
		return nil
	}

	userRepo := persistence.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo)

	adminRepo := persistence.NewAdminRepository(db)
	adminUsecase := usecase_admin.NewAdminUsecase(adminRepo)
	jwtHelper := utils.NewJWTHelper(conf, adminRepo)

	allUsecase := usecase.AllUsecase{
		AuthUsecase: authUsecase,
		Admin: usecase.AdminUsecase{
			AdminUsecase: adminUsecase,
		},
	}
	httpRest := rest.NewHttpServer(allUsecase, jwtHelper)
	httpRest.Run(conf)
	httpRest.GracefulShutdown = handlerShutdown
}
