package model

import "time"

// User はユーザーモデルを表します
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// IsValid はユーザーが有効かどうかを検証します
func (u *User) IsValid() bool {
	return u.Email != "" && u.Name != ""
}
