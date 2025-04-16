package util

import (
	"testing"

	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"min=18,max=120"`
}

func TestValidator_Validate(t *testing.T) {
	v := NewValidator()

	t.Run("Valid struct", func(t *testing.T) {
		valid := TestStruct{
			Name:  "Test User",
			Email: "test@example.com",
			Age:   30,
		}
		err := v.Validate(valid)
		assert.NoError(t, err)
	})

	t.Run("Missing required field", func(t *testing.T) {
		invalid := TestStruct{
			Email: "test@example.com",
			Age:   30,
		}
		err := v.Validate(invalid)
		assert.Error(t, err)

		appError, ok := err.(*model.AppError)
		assert.True(t, ok)
		assert.Equal(t, model.ErrInvalidInput, appError.Code, "エラーコードが一致しません")

	})

	t.Run("Invalid email", func(t *testing.T) {
		invalid := TestStruct{
			Name:  "Test User",
			Email: "not-an-email",
			Age:   30,
		}
		err := v.Validate(invalid)
		assert.Error(t, err)

		appError, ok := err.(*model.AppError)
		assert.True(t, ok)
		assert.Equal(t, model.ErrInvalidInput, appError.Code, "エラーコードが一致しません")
		assert.Contains(t, appError.Message, "emailは有効なメールアドレスである必要があります")
	})

	t.Run("Multiple validation errors", func(t *testing.T) {
		invalid := TestStruct{
			Age: 10, // Too young
		}
		err := v.Validate(invalid)
		assert.Error(t, err)

		appError, ok := err.(*model.AppError)
		assert.True(t, ok)
		assert.Equal(t, model.ErrInvalidInput, appError.Code, "エラーコードが一致しません")
		assert.Contains(t, appError.Message, "nameは必須です")
		assert.Contains(t, appError.Message, "emailは必須です")
		assert.Contains(t, appError.Message, "ageは18文字以上である必要があります")
	})
}
