package service

import (
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/repository"
)

// ProductService は製品関連のビジネスロジックを提供します
type ProductService struct {
	productRepo repository.ProductRepository
}

// NewProductService は新しいProductServiceインスタンスを作成します
func NewProductService(productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// GetProduct は指定されたIDの製品を取得します
func (s *ProductService) GetProduct(id int) (*model.Product, error) {
	return s.productRepo.GetByID(id)
}

// ListProducts は製品の一覧を取得します
func (s *ProductService) ListProducts(page, pageSize int) ([]*model.Product, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return s.productRepo.List(pageSize, offset)
}

// CreateProduct は新しい製品を作成します
func (s *ProductService) CreateProduct(product *model.Product) error {
	// 入力検証
	if !product.IsValid() {
		return model.NewInvalidInputError("無効な製品データです", nil)
	}

	// 製品作成
	return s.productRepo.Create(product)
}

// UpdateProduct は製品情報を更新します
func (s *ProductService) UpdateProduct(product *model.Product) error {
	// 入力検証
	if !product.IsValid() {
		return model.NewInvalidInputError("無効な製品データです", nil)
	}

	// 製品の存在確認
	_, err := s.productRepo.GetByID(product.ID)
	if err != nil {
		return err
	}

	// 製品更新
	return s.productRepo.Update(product)
}

// DeleteProduct は製品を削除します
func (s *ProductService) DeleteProduct(id int) error {
	// 製品の存在確認
	_, err := s.productRepo.GetByID(id)
	if err != nil {
		return err
	}

	// 製品削除
	return s.productRepo.Delete(id)
}
