package dto

import (
	"time"

	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
)

// ProductRequest は製品作成/更新リクエストのDTOです
type ProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

// ProductResponse は製品情報レスポンスのDTOです
type ProductResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToModel はDTOをモデルに変換します
func (r *ProductRequest) ToModel() *model.Product {
	return &model.Product{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
	}
}

// FromModel はモデルからDTOを作成します
func FromProductModel(product *model.Product) *ProductResponse {
	return &ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

// ProductsResponse は複数製品のレスポンスDTOです
type ProductsResponse struct {
	Products []ProductResponse `json:"products"`
	Total    int               `json:"total"`
	Page     int               `json:"page"`
	Size     int               `json:"size"`
}

// FromProductModels はモデルのスライスからDTOを作成します
func FromProductModels(products []*model.Product, page, size, total int) *ProductsResponse {
	response := &ProductsResponse{
		Products: make([]ProductResponse, len(products)),
		Total:    total,
		Page:     page,
		Size:     size,
	}

	for i, product := range products {
		response.Products[i] = *FromProductModel(product)
	}

	return response
}
