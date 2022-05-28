package db_postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"loginUserGo/infrastructure/database"
)

func InitGorm(config *database.Config) *gorm.DB {

	dsn := "host=" + config.Hostname + " user=" + config.UserName + " password=" + config.Password + " dbname=" +
		config.Database + " port=" + config.Port + " sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic("failed to connect postgres database")
	}

	return db
}
