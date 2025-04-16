package model

import (
	"errors"
	"fmt"
)

// ErrorCode はアプリケーションエラーコードを表します
type ErrorCode string

const (
	// エラーコード定義
	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrInvalidInput  ErrorCode = "INVALID_INPUT"
	ErrUnauthorized  ErrorCode = "UNAUTHORIZED"
	ErrBadRequest    ErrorCode = "BAD_REQUEST"
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

// IsNotFoundError はエラーが「見つからない」エラーかどうかを判定します
func IsNotFoundError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Code == ErrNotFound
}

// IsInvalidInputError はエラーが「無効な入力」エラーかどうかを判定します
func IsInvalidInputError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Code == ErrInvalidInput
}

// IsUnauthorizedError はエラーが「認証エラー」かどうかを判定します
func IsUnauthorizedError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Code == ErrUnauthorized
}

// IsInternalError はエラーが「内部エラー」かどうかを判定します
func IsInternalError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Code == ErrInternalError
}

// IsDuplicateError はエラーが「重複エラー」かどうかを判定します
func IsDuplicateError(err error) bool {
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Code == ErrDuplicate
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
