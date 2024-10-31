package database

import (
	"fmt"
	"github.com/auperman-lab/lab2/internal/configs"
	"github.com/auperman-lab/lab2/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func LoadDatabase() *gorm.DB {
	db := connect()
	err := db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.SubCategory{},
		&models.DiscountPeriod{},
		&models.Distributor{})
	if err != nil {
		slog.Error("failed migrating database", "error", err)
	}
	return db
}

func connect() *gorm.DB {
	var err error

	dsn := fmt.Sprintf(`host=%s
	dbname=%s
	user=%s
	password=%s
	port=%d`,
		configs.Env.DBHost,
		configs.Env.DBName,
		configs.Env.DBUser,
		configs.Env.DBPassword,
		configs.Env.DBPort,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		slog.Error("failed connecting orm to database", "error", err)
	} else {
		slog.Info("Successfully connected to the Postgres database")
	}

	return database
}
