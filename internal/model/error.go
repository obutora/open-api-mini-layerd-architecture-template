package model

import "fmt"

// ErrorCode はアプリケーションエラーコードを表します
type ErrorCode string

const (
	// エラーコード定義
	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrInvalidInput  ErrorCode = "INVALID_INPUT"
	ErrUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrInternalError ErrorCode = "INTERNAL_ERROR"
	ErrDuplicate     ErrorCode = "DUPLICATE"
)

// AppError はアプリケーションエラーを表します
type AppError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"-"` // 内部エラー（JSONには含まれない）
}

// Error はエラーインターフェースを実装します
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%s)", e.Code, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Unwrap は内部エラーを返します
func (e *AppError) Unwrap() error {
	return e.Err
}

// NewNotFoundError は「見つからない」エラーを作成します
func NewNotFoundError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrNotFound,
		Message: message,
		Err:     err,
	}
}

// NewInvalidInputError は「無効な入力」エラーを作成します
func NewInvalidInputError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrInvalidInput,
		Message: message,
		Err:     err,
	}
}

// NewUnauthorizedError は「認証エラー」を作成します
func NewUnauthorizedError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrUnauthorized,
		Message: message,
		Err:     err,
	}
}

// NewInternalError は「内部エラー」を作成します
func NewInternalError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrInternalError,
		Message: message,
		Err:     err,
	}
}

// NewDuplicateError は「重複エラー」を作成します
func NewDuplicateError(message string, err error) *AppError {
	return &AppError{
		Code:    ErrDuplicate,
		Message: message,
		Err:     err,
	}
}
