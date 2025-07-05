package services

import (
	"errors"
	"rest-api/internal/models"
	"rest-api/internal/repositories"
	"rest-api/internal/utils"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	productRepo *repositories.ProductRepository
	validator   *validator.Validate
}

func NewProductService(productRepo *repositories.ProductRepository, validator *validator.Validate) *ProductService {
	return &ProductService{
		productRepo: productRepo,
		validator:   validator,
	}
}

func (s *ProductService) CreateProduct(userID string, req *models.CreateProductRequest) (*models.ProductResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid user ID")
	}

	product := &models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		UserID:      objID,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	// Load user data
	product, err = s.productRepo.GetByID(product.ID.Hex())
	if err != nil {
		return nil, err
	}

	return s.convertToProductResponse(product), nil
}

func (s *ProductService) GetProductByID(id string) (*models.ProductResponse, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return s.convertToProductResponse(product), nil
}

func (s *ProductService) GetAllProducts(page, limit int) (*models.PaginatedResponse, error) {
	page, limit = utils.GetPaginationParams(page, limit)
	offset := utils.CalculateOffset(page, limit)

	products, total, err := s.productRepo.GetAll(offset, limit)
	if err != nil {
		return nil, err
	}

	productResponses := make([]models.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = *s.convertToProductResponse(&product)
	}

	return &models.PaginatedResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    productResponses,
		Page:    page,
		Limit:   limit,
		Total:   total,
	}, nil
}

func (s *ProductService) GetProductsByUserID(userID string, page, limit int) (*models.PaginatedResponse, error) {
	page, limit = utils.GetPaginationParams(page, limit)
	offset := utils.CalculateOffset(page, limit)

	products, total, err := s.productRepo.GetByUserID(userID, offset, limit)
	if err != nil {
		return nil, err
	}

	productResponses := make([]models.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = *s.convertToProductResponse(&product)
	}

	return &models.PaginatedResponse{
		Success: true,
		Message: "Products retrieved successfully",
		Data:    productResponses,
		Page:    page,
		Limit:   limit,
		Total:   total,
	}, nil
}

func (s *ProductService) UpdateProduct(id, userID string, req *models.UpdateProductRequest) (*models.ProductResponse, error) {
	if err := s.validator.Struct(req); err != nil {
		return nil, err
	}

	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if product.UserID.Hex() != userID {
		return nil, errors.New("you can only update your own products")
	}

	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price > 0 {
		product.Price = req.Price
	}
	if req.Stock >= 0 {
		product.Stock = req.Stock
	}

	if err := s.productRepo.Update(product); err != nil {
		return nil, err
	}

	return s.convertToProductResponse(product), nil
}

func (s *ProductService) DeleteProduct(id, userID string) error {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	if product.UserID.Hex() != userID {
		return errors.New("you can only delete your own products")
	}

	return s.productRepo.Delete(id)
}

func (s *ProductService) convertToProductResponse(product *models.Product) *models.ProductResponse {
	return &models.ProductResponse{
		ID:          product.ID.Hex(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		User: models.UserResponse{
			ID:        product.User.ID.Hex(),
			Name:      product.User.Name,
			Email:     product.User.Email,
			CreatedAt: product.User.CreatedAt.Format(time.RFC3339),
			UpdatedAt: product.User.UpdatedAt.Format(time.RFC3339),
		},
		CreatedAt: product.CreatedAt.Format(time.RFC3339),
		UpdatedAt: product.UpdatedAt.Format(time.RFC3339),
	}
}
