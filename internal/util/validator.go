package util

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
)

// Validator はアプリケーションのバリデーターを表します
type Validator struct {
	validator *validator.Validate
}

// NewValidator は新しいValidatorインスタンスを作成します
func NewValidator() *Validator {
	v := validator.New()

	// JSONタグをフィールド名として使用
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// カスタムバリデーションの登録
	_ = v.RegisterValidation("phone", validatePhone)
	_ = v.RegisterValidation("password", validatePassword)

	return &Validator{
		validator: v,
	}
}

// Validate は構造体のバリデーションを行います
func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		// バリデーションエラーをアプリケーションエラーに変換
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			messages := make([]string, 0, len(validationErrors))
			for _, e := range validationErrors {
				messages = append(messages, formatValidationError(e))
			}
			return model.NewInvalidInputError(strings.Join(messages, "; "), err)
		}
		return model.NewInvalidInputError("入力データが無効です", err)
	}
	return nil
}

// ValidateVar は単一の変数のバリデーションを行います
func (v *Validator) ValidateVar(field interface{}, tag string) error {
	if err := v.validator.Var(field, tag); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			messages := make([]string, 0, len(validationErrors))
			for _, e := range validationErrors {
				messages = append(messages, formatValidationError(e))
			}
			return model.NewInvalidInputError(strings.Join(messages, "; "), err)
		}
		return model.NewInvalidInputError("入力データが無効です", err)
	}
	return nil
}

// formatValidationError はバリデーションエラーを人間が読みやすい形式にフォーマットします
func formatValidationError(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + "は必須です"
	case "email":
		return e.Field() + "は有効なメールアドレスである必要があります"
	case "min":
		return e.Field() + "は" + e.Param() + "文字以上である必要があります"
	case "max":
		return e.Field() + "は" + e.Param() + "文字以下である必要があります"
	case "phone":
		return e.Field() + "は有効な電話番号である必要があります"
	case "password":
		return e.Field() + "は8文字以上で、少なくとも1つの数字、1つの大文字、1つの小文字、1つの特殊文字を含む必要があります"
	default:
		return e.Field() + "は" + e.Tag() + "のルールに違反しています"
	}
}

// validatePhone は電話番号のバリデーションを行います
func validatePhone(fl validator.FieldLevel) bool {
	// 日本の電話番号形式（例: 090-1234-5678）
	re := regexp.MustCompile(`^(0\d{1,4}-\d{1,4}-\d{4})$`)
	return re.MatchString(fl.Field().String())
}

// validatePassword はパスワードのバリデーションを行います
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	// 8文字以上
	if len(password) < 8 {
		return false
	}

	// 少なくとも1つの数字
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	// 少なくとも1つの大文字
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	// 少なくとも1つの小文字
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	// 少なくとも1つの特殊文字
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasNumber && hasUpper && hasLower && hasSpecial
}
