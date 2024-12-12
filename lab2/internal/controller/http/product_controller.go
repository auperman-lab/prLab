package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/auperman-lab/lab2/internal/models"
	"github.com/auperman-lab/lab2/internal/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
)

type IProductService interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProductByID(ctx context.Context, id uint) (*models.ReturnProduct, *models.Image, error)
	GetProductByName(ctx context.Context, name string) (*models.ReturnProduct, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProductByID(ctx context.Context, id uint) error
	GetAllProducts(ctx context.Context, pag utils.Pagination) ([]models.Product, error)
	UpdateProductImage(ctx context.Context, img []byte, id uint) error
}

type ProductController struct {
	productService IProductService
}

func NewProductController(service IProductService) *ProductController {
	slog.Info("Creating new product controller")
	return &ProductController{
		productService: service,
	}
}

func (ctrl *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var createProduct models.CreateProduct
	if err := utils.ParseJSON(r, &createProduct); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	subCategoryID, exists := models.FindSubCategory(createProduct.SubCategory)
	if !exists {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid subcategory: %s", createProduct.SubCategory))
		return
	}

	product := models.Product{
		DistributorID:    createProduct.DistributorID,
		Name:             createProduct.Name,
		Price:            createProduct.Price,
		PriceOld:         createProduct.PriceOld,
		Discount:         createProduct.Discount,
		DiscountPeriodID: createProduct.DiscountPeriodID,
		Available:        createProduct.Available,
		Link:             createProduct.Link,
		SubCategoryID:    subCategoryID.ID,
		SpecialCondition: createProduct.SpecialCondition,
		ImageID:          createProduct.ImageID,
	}

	ctx := r.Context()

	if err := ctrl.productService.CreateProduct(ctx, &product); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/products/%d", product.ID))
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Product created successfully"})
}
func (ctrl *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing id"))
	}

	id, err := strconv.ParseUint(idStr, 10, 32) // Parsing as uint32 for GORM
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	ctx := r.Context()
	product, img, err := ctrl.productService.GetProductByID(ctx, uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(w, http.StatusNotFound, errors.New("product not found"))
			return
		}
		http.Error(w, "failed to get product", http.StatusInternalServerError)
		return
	}

	//----------------------------------------------------

	mw := multipart.NewWriter(w)
	w.Header().Set("Content-Type", mw.FormDataContentType())
	defer mw.Close()

	productJSON, err := json.Marshal(product)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to marshal product data"))
		return
	}

	partJSON, err := mw.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{`form-data; name="product_data"`},
		"Content-Type":        []string{"application/json"},
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to create JSON part"))
		return
	}
	_, err = partJSON.Write(productJSON)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to write JSON part"))
		return
	}

	if img != nil {
		partImage, err := mw.CreatePart(textproto.MIMEHeader{
			"Content-Disposition": []string{`form-data; name="image"; filename="product_image.png"`},
			"Content-Type":        []string{"image/png"},
		})
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to create image part"))
			return
		}
		_, err = partImage.Write(img.Image)
		if err != nil {
			utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to write image part"))
			return
		}
	}

	mw.Close()

}
func (ctrl *ProductController) GetProductByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nameStr, ok := vars["name"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing id"))
	}

	ctx := r.Context()
	product, err := ctrl.productService.GetProductByName(ctx, nameStr)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(w, http.StatusNotFound, errors.New("product not found"))
			return
		}
		http.Error(w, "failed to get product", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}
func (ctrl *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := utils.ParseJSON(r, &product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	defer r.Body.Close()

	ctx := r.Context()

	if err := ctrl.productService.UpdateProduct(ctx, &product); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("/products/%d", product.ID))
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "Product updated successfully"})
}
func (ctrl *ProductController) DeleteProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing id"))
	}

	id, err := strconv.ParseUint(idStr, 10, 32) // Parsing as uint32 for GORM
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	ctx := r.Context()
	if err := ctrl.productService.DeleteProductByID(ctx, uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(w, http.StatusNotFound, errors.New("product not found"))
			return
		}
		http.Error(w, "failed to get product", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}
func (ctrl *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pageStr, ok := vars["page"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing page parameter"))
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid page parameter"))
		return
	}

	limitStr, ok := vars["limit"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing limit parameter"))
	}
	limit, err := strconv.Atoi(limitStr) // Parsing as uint32 for GORM
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid limit parameter"))
		return
	}
	pag := utils.Pagination{
		Page:  page,
		Limit: limit,
	}

	ctx := r.Context()
	product, err := ctrl.productService.GetAllProducts(ctx, pag)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.WriteError(w, http.StatusNotFound, errors.New("products not found"))
			return
		}
		http.Error(w, "failed to get product", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, product)
}
func (ctrl *ProductController) UpdateProductImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, errors.New("missing id"))
	}
	id, err := strconv.ParseUint(idStr, 10, 32) // Parsing as uint32 for GORM
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}
	fmt.Println(id)

	if err := r.ParseMultipartForm(1 << 20); err != nil {
		slog.Error("Failed to parse multipart form", "error", err)
		http.Error(w, "Failed to parse multipart form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		slog.Error("Failed to get file from form data", "error", err)
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Failed to read file data", "error", err)
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}
	ctx := r.Context()
	if err := ctrl.productService.UpdateProductImage(ctx, imageData, uint(id)); err != nil {
		slog.Error("Failed to update product image", "error", err)
		http.Error(w, "Failed to update product image", http.StatusInternalServerError)
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Image uploaded successfully", "product_id": idStr})
}
