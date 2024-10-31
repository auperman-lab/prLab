package service

import (
	"context"
	"github.com/auperman-lab/lab2/internal/models"
	"github.com/auperman-lab/lab2/internal/utils"
	"log/slog"
)

type IProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProductByID(ctx context.Context, id uint) (*models.Product, error)
	GetProductByName(ctx context.Context, name string) (*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProductByID(ctx context.Context, id uint) error
	GetAllProducts(ctx context.Context, pag utils.Pagination) ([]models.Product, error)
	UpdateProductImage(ctx context.Context, img []byte, id uint) error
}

type ProductService struct {
	productRepository IProductRepository
}

func NewProductService(repo IProductRepository) *ProductService {
	slog.Info("Creating new product service")
	return &ProductService{productRepository: repo}
}

func (svc *ProductService) CreateProduct(ctx context.Context, product *models.Product) error {
	slog.Info("Creating product...", "name", product.Name)
	return svc.productRepository.CreateProduct(ctx, product)
}
func (svc *ProductService) GetProductByID(ctx context.Context, id uint) (*models.Product, error) {
	slog.Info("Getting product...", "id", id)
	return svc.productRepository.GetProductByID(ctx, id)
}
func (svc *ProductService) GetProductByName(ctx context.Context, name string) (*models.Product, error) {
	slog.Info("Getting product...", "name", name)
	return svc.productRepository.GetProductByName(ctx, name)
}
func (svc *ProductService) UpdateProduct(ctx context.Context, product *models.Product) error {
	slog.Info("Updating product...", "name", product.Name)
	return svc.productRepository.UpdateProduct(ctx, product)
}
func (svc *ProductService) DeleteProductByID(ctx context.Context, id uint) error {
	slog.Info("Deleting product...", "id", id)
	return svc.productRepository.DeleteProductByID(ctx, id)
}
func (svc *ProductService) GetAllProducts(ctx context.Context, pag utils.Pagination) ([]models.Product, error) {
	slog.Info("Getting all products...")
	return svc.productRepository.GetAllProducts(ctx, pag)
}
func (svc *ProductService) UpdateProductImage(ctx context.Context, img []byte, id uint) error {
	slog.Info("Updating product image...", "id", id)
	return svc.productRepository.UpdateProductImage(ctx, img, id)
}
