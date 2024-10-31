package repository

import (
	"context"
	"github.com/auperman-lab/lab2/internal/models"
	"github.com/auperman-lab/lab2/internal/utils"
	"gorm.io/gorm"
	"log/slog"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	slog.Info("Creating new product repository")
	return &ProductRepository{
		db: db,
	}
}

func (repo *ProductRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	err := repo.db.WithContext(ctx).Create(&product).Error
	if err != nil {
		slog.Error("Failed to create a product", "error", err.Error())
		return err
	}
	return nil
}
func (repo *ProductRepository) GetProductByID(ctx context.Context, id uint) (*models.Product, error) {
	var product = &models.Product{}
	err := repo.db.WithContext(ctx).Where("id=?", id).First(&product).Error
	if err != nil {
		slog.Error("Failed to find product by id", "error", err.Error())
		return nil, err
	}
	return product, nil
}
func (repo *ProductRepository) GetProductByName(ctx context.Context, name string) (*models.Product, error) {
	var product = &models.Product{}
	err := repo.db.WithContext(ctx).Where("name=?", name).First(&product).Error
	if err != nil {
		slog.Error("Failed to find product by name", "error", err.Error())
		return nil, err
	}
	return product, nil
}
func (repo *ProductRepository) UpdateProduct(ctx context.Context, product *models.Product) error {
	var existingProduct = &models.Product{}
	err := repo.db.WithContext(ctx).Where("id=?", product.ID).First(&existingProduct).Error
	if err != nil {
		slog.Error("Failed to find petition to update", "error", err.Error(), "id", product.ID)
		return err
	}
	slog.Info("Updated product", "id", product.ID)

	if product.Name != "" {
		slog.Info("Updating product", "name", product.Name)
		existingProduct.Name = product.Name
	}
	if product.Price != 0 {
		slog.Info("Updating product", "price", product.Price)
		existingProduct.Price = product.Price
	}
	if product.PriceOld != 0 {
		slog.Info("Updating product", "priceOld", product.PriceOld)
		existingProduct.PriceOld = product.PriceOld
	}
	if product.Discount != 0 {
		slog.Info("Updating product", "discount", product.Discount)
		existingProduct.Discount = product.Discount
	}
	if product.Available != nil {
		slog.Info("Updating product", "available", product.Available)
		existingProduct.Available = product.Available
	}
	if product.Link != "" {
		slog.Info("Updating product", "link", product.Link)
		existingProduct.Link = product.Link
	}
	if product.ImageID != nil {
		slog.Info("Updating product", "image", product.ImageID)
		existingProduct.ImageID = product.ImageID
	}
	if product.SpecialCondition != "" {
		slog.Info("Updating product", "specialCondition", product.SpecialCondition)
		existingProduct.SpecialCondition = product.SpecialCondition
	}
	if product.DiscountPeriodID != nil {
		slog.Info("Updating product", "discountPeriodId", product.DiscountPeriodID)
		existingProduct.DiscountPeriodID = product.DiscountPeriodID
	}
	if product.SubCategoryID != 0 {
		slog.Info("Updating product", "subCategoryId", product.SubCategoryID)
		existingProduct.SubCategoryID = product.SubCategoryID
	}

	if err := repo.db.WithContext(ctx).Save(&existingProduct).Error; err != nil {
		slog.Error("Failed to update product", "error", err.Error())
		return err
	}

	return nil
}
func (repo *ProductRepository) DeleteProductByID(ctx context.Context, id uint) error {
	err := repo.db.WithContext(ctx).Where("id=?", id).Delete(&models.Product{}).Error
	if err != nil {
		slog.Error("Failed to delete product", "error", err.Error())
		return err
	}
	return nil
}
func (repo *ProductRepository) GetAllProducts(ctx context.Context, pag utils.Pagination) ([]models.Product, error) {
	var products []models.Product
	err := repo.db.WithContext(ctx).Offset(pag.Page).Limit(pag.Limit).Find(&products).Error
	if err != nil {
		slog.Error("Failed to find all products", "error", err.Error())
		return nil, err
	}
	return products, nil
}
func (repo *ProductRepository) UpdateProductImage(ctx context.Context, img []byte, id uint) error {
	var existingProduct = &models.Product{}

	if err := repo.db.WithContext(ctx).Where("id=?", id).First(&existingProduct).Error; err != nil {
		slog.Error("Failed to find petition to update", "error", err.Error())
		return err
	}

	newImage := models.Image{Image: img}
	if err := repo.db.WithContext(ctx).Model(&models.Image{}).Create(&newImage).Error; err != nil {
		slog.Error("Failed to update product", "error", err.Error())
		return err
	}

	existingProduct.ImageID = &newImage.ID

	if err := repo.db.WithContext(ctx).Save(&existingProduct).Error; err != nil {
		slog.Error("Failed to update product image", "error", err.Error())
		return err
	}

	return nil
}
