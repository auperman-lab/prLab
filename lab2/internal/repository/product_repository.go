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
func (repo *ProductRepository) GetProductByID(ctx context.Context, id uint) (*models.ReturnProduct, error) {
	var product models.Product
	err := repo.db.WithContext(ctx).Where("id=?", id).First(&product).Error
	if err != nil {
		slog.Error("Failed to find product by id", "error", err.Error())
		return nil, err
	}

	returnProduct := &models.ReturnProduct{
		ID:               product.ID,
		Name:             product.Name,
		Price:            product.Price,
		PriceOld:         product.PriceOld,
		Discount:         product.Discount,
		Available:        product.Available,
		Link:             product.Link,
		ImageID:          product.ImageID,
		SpecialCondition: product.SpecialCondition,
	}

	var distributor models.Distributor
	if err := repo.db.WithContext(ctx).Where("id=?", product.DistributorID).First(&distributor).Error; err == nil {
		returnProduct.Distributor = distributor.Name
	}

	subCategory, err := repo.GetSubCategoryByID(ctx, product.SubCategoryID)
	if err == nil {
		returnProduct.SubCategory = subCategory.Name
	}

	category, err := repo.GetCategoryByID(ctx, subCategory.CategoryID)
	if err == nil {
		returnProduct.Category = category.Name
	}
	var discountPeriod models.DiscountPeriod
	if err := repo.db.WithContext(ctx).Where("id=?", product.DiscountPeriodID).First(&discountPeriod).Error; err == nil {
		returnProduct.DiscountPeriodStart = discountPeriod.StartDate
		returnProduct.DiscountPeriodEnd = discountPeriod.EndDate
	}

	return returnProduct, nil
}

func (repo *ProductRepository) GetProductByName(ctx context.Context, name string) (*models.ReturnProduct, error) {
	var product = &models.Product{}
	err := repo.db.WithContext(ctx).Where("name=?", name).First(&product).Error
	if err != nil {
		slog.Error("Failed to find product by name", "error", err.Error())
		return nil, err
	}

	returnProduct := &models.ReturnProduct{
		ID:               product.ID,
		Name:             product.Name,
		Price:            product.Price,
		PriceOld:         product.PriceOld,
		Discount:         product.Discount,
		Available:        product.Available,
		Link:             product.Link,
		ImageID:          product.ImageID,
		SpecialCondition: product.SpecialCondition,
	}

	var distributor models.Distributor
	if err := repo.db.WithContext(ctx).Where("id=?", product.DistributorID).First(&distributor).Error; err == nil {
		returnProduct.Distributor = distributor.Name
	}

	subCategory, err := repo.GetSubCategoryByID(ctx, product.SubCategoryID)
	if err == nil {
		returnProduct.SubCategory = subCategory.Name
	}

	category, err := repo.GetCategoryByID(ctx, subCategory.CategoryID)
	if err == nil {
		returnProduct.Category = category.Name
	}

	var discountPeriod models.DiscountPeriod
	if err := repo.db.WithContext(ctx).Where("id=?", product.DiscountPeriodID).First(&discountPeriod).Error; err == nil {
		returnProduct.DiscountPeriodStart = discountPeriod.StartDate
		returnProduct.DiscountPeriodEnd = discountPeriod.EndDate
	}

	return returnProduct, nil
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
	if product.ImageID != nil && *product.ImageID != 0 {
		slog.Info("Updating product", "image", product.ImageID)
		existingProduct.ImageID = product.ImageID
	}
	if product.SpecialCondition != "" {
		slog.Info("Updating product", "specialCondition", product.SpecialCondition)
		existingProduct.SpecialCondition = product.SpecialCondition
	}
	if product.DiscountPeriodID != nil && *product.DiscountPeriodID != 0 {
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

	newImage := models.Image{Image: img}
	if err := repo.db.WithContext(ctx).Model(&models.Image{}).Create(&newImage).Error; err != nil {
		slog.Error("Failed to update product", "error", err.Error())
		return err
	}

	var product = &models.Product{}
	product.ID = id
	product.ImageID = &newImage.ID

	err := repo.UpdateProduct(ctx, product)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) GetCategoryByID(ctx context.Context, id uint) (*models.Category, error) {
	var category = &models.Category{}
	if err := repo.db.WithContext(ctx).Where("id=?", id).First(&category).Error; err != nil {
		slog.Error("Failed to find category", "error", err.Error())
		return nil, err
	}
	return category, nil
}
func (repo *ProductRepository) GetSubCategoryByID(ctx context.Context, id uint) (*models.SubCategory, error) {
	var subCategory = &models.SubCategory{}
	if err := repo.db.WithContext(ctx).Where("id=?", id).First(&subCategory).Error; err != nil {
		slog.Error("Failed to find subcategory", "error", err.Error())
		return nil, err
	}
	return subCategory, nil
}
