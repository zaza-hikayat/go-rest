package persistence

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/zaza-hikayat/go-rest-sample/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(conf config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		conf.DB.Host,
		conf.DB.Port,
		conf.DB.Username,
		conf.DB.Password,
		conf.DB.DBName,
		conf.DB.SSLMode,
		conf.DB.Timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatal("Failed to connect to database:", err)
	}

	return db
}
