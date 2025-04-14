package repository

import (
	"database/sql"

	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
)

// UserRepository はユーザーデータアクセスのインターフェースを定義します
type UserRepository interface {
	GetByID(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	List(limit, offset int) ([]*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id int) error
}

// UserRepositoryImpl はUserRepositoryインターフェースの実装です
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository は新しいUserRepositoryインスタンスを作成します
func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// GetByID は指定されたIDのユーザーを取得します
func (r *UserRepositoryImpl) GetByID(id int) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = ?`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.NewNotFoundError("ユーザーが見つかりません", err)
		}
		return nil, model.NewInternalError("ユーザー取得中にエラーが発生しました", err)
	}

	return user, nil
}

// GetByEmail は指定されたメールアドレスのユーザーを取得します
func (r *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE email = ?`

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.NewNotFoundError("ユーザーが見つかりません", err)
		}
		return nil, model.NewInternalError("ユーザー取得中にエラーが発生しました", err)
	}

	return user, nil
}

// List はユーザーの一覧を取得します
func (r *UserRepositoryImpl) List(limit, offset int) ([]*model.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users LIMIT ? OFFSET ?`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, model.NewInternalError("ユーザー一覧取得中にエラーが発生しました", err)
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, model.NewInternalError("ユーザーデータ読み取り中にエラーが発生しました", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, model.NewInternalError("ユーザーデータ読み取り中にエラーが発生しました", err)
	}

	return users, nil
}

// Create は新しいユーザーを作成します
func (r *UserRepositoryImpl) Create(user *model.User) error {
	query := `INSERT INTO users (name, email, created_at, updated_at) VALUES (?, ?, NOW(), NOW())`

	result, err := r.db.Exec(query, user.Name, user.Email)
	if err != nil {
		return model.NewInternalError("ユーザー作成中にエラーが発生しました", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.NewInternalError("ユーザーID取得中にエラーが発生しました", err)
	}

	user.ID = int(id)
	return nil
}

// Update はユーザー情報を更新します
func (r *UserRepositoryImpl) Update(user *model.User) error {
	query := `UPDATE users SET name = ?, email = ?, updated_at = NOW() WHERE id = ?`

	result, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return model.NewInternalError("ユーザー更新中にエラーが発生しました", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.NewInternalError("更新結果確認中にエラーが発生しました", err)
	}

	if rowsAffected == 0 {
		return model.NewNotFoundError("更新対象のユーザーが見つかりません", nil)
	}

	return nil
}

// Delete はユーザーを削除します
func (r *UserRepositoryImpl) Delete(id int) error {
	query := `DELETE FROM users WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return model.NewInternalError("ユーザー削除中にエラーが発生しました", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return model.NewInternalError("削除結果確認中にエラーが発生しました", err)
	}

	if rowsAffected == 0 {
		return model.NewNotFoundError("削除対象のユーザーが見つかりません", nil)
	}

	return nil
}
