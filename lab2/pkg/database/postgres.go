package database

import (
	"fmt"
	"github.com/auperman-lab/lab2/internal/configs"
	"github.com/auperman-lab/lab2/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"log/slog"
)

func LoadDatabase() *gorm.DB {
	db := connect()
	err := db.AutoMigrate(
		&models.Product{},
		&models.Category{},
		&models.SubCategory{},
		&models.DiscountPeriod{},
		&models.Distributor{},
		&models.Image{})
	if err != nil {
		slog.Error("failed migrating database", "error", err)
		log.Fatal()
	}
	seedCategory(db)
	seedDistributor(db)
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
		log.Fatal()
	} else {
		slog.Info("Successfully connected to the Postgres database")
	}

	return database
}

func seedCategory(db *gorm.DB) {
	var count int64
	db.Model(&models.Category{}).Count(&count)
	if count == 0 {
		for _, category := range models.CategorySeedData {
			db.Create(&category)
		}
	}
	count = 0
	db.Model(&models.SubCategory{}).Count(&count)
	if count == 0 {
		for _, subCategory := range models.SubCategorySeedData {
			db.Create(&subCategory)
		}
	}
}
func seedDistributor(db *gorm.DB) {
	var count int64
	db.Model(&models.Distributor{}).Count(&count)
	if count == 0 {
		for _, distributor := range models.DistributorSeedData {
			db.Create(&distributor)
		}
	}
}
