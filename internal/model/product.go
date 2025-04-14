package model

import "time"

// Product は製品モデルを表します
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// IsValid は製品が有効かどうかを検証します
func (p *Product) IsValid() bool {
	return p.Name != "" && p.Price > 0
}
