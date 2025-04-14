package service

import (
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/repository"
)

// UserService はユーザー関連のビジネスロジックを提供します
type UserService struct {
	userRepo repository.UserRepository
}

// NewUserService は新しいUserServiceインスタンスを作成します
func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// GetUser は指定されたIDのユーザーを取得します
func (s *UserService) GetUser(id int) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

// GetUserByEmail は指定されたメールアドレスのユーザーを取得します
func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.userRepo.GetByEmail(email)
}

// ListUsers はユーザーの一覧を取得します
func (s *UserService) ListUsers(page, pageSize int) ([]*model.User, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	return s.userRepo.List(pageSize, offset)
}

// CreateUser は新しいユーザーを作成します
func (s *UserService) CreateUser(user *model.User) error {
	// 入力検証
	if !user.IsValid() {
		return model.NewInvalidInputError("無効なユーザーデータです", nil)
	}

	// メールアドレスの重複チェック
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err == nil && existingUser != nil {
		return model.NewDuplicateError("このメールアドレスは既に使用されています", nil)
	}

	// ユーザー作成
	return s.userRepo.Create(user)
}

// UpdateUser はユーザー情報を更新します
func (s *UserService) UpdateUser(user *model.User) error {
	// 入力検証
	if !user.IsValid() {
		return model.NewInvalidInputError("無効なユーザーデータです", nil)
	}

	// ユーザーの存在確認
	_, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}

	// メールアドレスの重複チェック（自分以外）
	existingUser, err := s.userRepo.GetByEmail(user.Email)
	if err == nil && existingUser != nil && existingUser.ID != user.ID {
		return model.NewDuplicateError("このメールアドレスは既に使用されています", nil)
	}

	// ユーザー更新
	return s.userRepo.Update(user)
}

// DeleteUser はユーザーを削除します
func (s *UserService) DeleteUser(id int) error {
	// ユーザーの存在確認
	_, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	// ユーザー削除
	return s.userRepo.Delete(id)
}
