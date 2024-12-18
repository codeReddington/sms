package database

import (
	"go.uber.org/zap"
	"go_fiber/internal/app/model"
	"go_fiber/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func Connect() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal("unable to connect to database :: ", zap.Error(err))
		panic(err)
	}
	Db = db
	return db
}

func Migrate() {

	migrations := Migrations{
		DB: Db,
		Models: []interface{}{
			&model.Task{},
		},
	}
	RunMigrations(migrations)
}

type Migrations struct {
	DB     *gorm.DB
	Models []interface{}
}

func RunMigrations(migrations Migrations) {
	for _, dbModels := range migrations.Models {
		err := migrations.DB.AutoMigrate(dbModels)
		if err != nil {
			logger.Error("Could not migrate %s", zap.Error(err))

		}
	}
}
