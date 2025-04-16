package dto

import (
	"time"

	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
)

// UserRequest はユーザー作成/更新リクエストのDTOです
type UserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"requiredemail"`
}

// CreateUserRequest はユーザー作成リクエストのDTOです
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"requiredemail"`
}

// UpdateUserRequest はユーザー更新リクエストのDTOです
type UpdateUserRequest struct {
	Name  string `json:"name" binding:"omitempty"`
	Email string `json:"email" binding:"omitemptyemail"`
}

// UserResponse はユーザー情報レスポンスのDTOです
type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToModel はDTOをモデルに変換します
func (r *UserRequest) ToModel() *model.User {
	return &model.User{
		Name:  r.Name,
		Email: r.Email,
	}
}

// FromModel はモデルからDTOを作成します
func FromUserModel(user *model.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// UsersResponse は複数ユーザーのレスポンスDTOです
type UsersResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
	Page  int            `json:"page"`
	Size  int            `json:"size"`
}

// FromUserModels はモデルのスライスからDTOを作成します
func FromUserModels(users []*model.User, page, size, total int) *UsersResponse {
	response := &UsersResponse{
		Users: make([]UserResponse, len(users)),
		Total: total,
		Page:  page,
		Size:  size,
	}

	for i, user := range users {
		response.Users[i] = *FromUserModel(user)
	}

	return response
}
