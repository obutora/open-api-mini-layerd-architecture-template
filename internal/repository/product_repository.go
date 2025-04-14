package repository

import (
	"database/sql"

	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
)

// ProductRepository は製品データアクセスのインターフェースを定義します
type ProductRepository interface {
	GetByID(id int) (*model.Product, error)
	List(limit, offset int) ([]*model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id int) error
}

// ProductRepositoryImpl はProductRepositoryインターフェースの実装です
type ProductRepositoryImpl struct {
	db *sql.DB
}

// NewProductRepository は新しいProductRepositoryインスタンスを作成します
func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

// GetByID は指定されたIDの製品を取得します
func (r *ProductRepositoryImpl) GetByID(id int) (*model.Product, error) {
	product := &model.Product{}
	query := `SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = ?`

	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.NewNotFoundError("製品が見つかりません", err)
		}
		return nil, model.NewInternalError("製品取得中にエラーが発生しました", err)
	}

	return product, nil
}

// List は製品の一覧を取得します
func (r *ProductRepositoryImpl) List(limit, offset int) ([]*model.Product, error) {
	query := `SELECT id, name, description, price, created_at, updated_at FROM products LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, model.NewInternalError("製品一覧取得中にエラーが発生しました", err)
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, model.NewInternalError("製品データ読み取り中にエラーが発生しました", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, model.NewInternalError("製品データ読み取り中にエラーが発生しました", err)
	}

	return products, nil
}

// Create は新しい製品を作成します
func (r *ProductRepositoryImpl) Create(product *model.Product) error {
	query := `INSERT INTO products (name, description, price, created_at, updated_at) VALUES (?, ?, ?, NOW(), NOW())`

	result, err := r.db.Exec(query, product.Name, product.Description, product.Price)
	if err != nil {
		return model.NewInternalError("製品作成中にエラーが発生しました", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.NewInternalError("製品ID取得中にエラーが発生しました", err)
	}

	product.ID = int(id)
	return nil
}

// Update は製品情報を更新します
func (r *ProductRepositoryImpl) Update(product *model.Product) error {
	query := `UPDATE products SET name = ?, description = ?, price = ?, updated_at = NOW() WHERE id = ?`

	result, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		return model.NewInternalError("製品更新中にエラーが発生しました", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.NewInternalError("更新結果確認中にエラーが発生しました", err)
	}

	if rowsAffected == 0 {
		return model.NewNotFoundError("更新対象の製品が見つかりません", nil)
	}

	return nil
}

// Delete は製品を削除します
func (r *ProductRepositoryImpl) Delete(id int) error {
	query := `DELETE FROM products WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return model.NewInternalError("製品削除中にエラーが発生しました", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.NewInternalError("削除結果確認中にエラーが発生しました", err)
	}

	if rowsAffected == 0 {
		return model.NewNotFoundError("削除対象の製品が見つかりません", nil)
	}

	return nil
}
