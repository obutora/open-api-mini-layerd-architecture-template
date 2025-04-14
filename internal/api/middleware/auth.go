package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/config"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/model"
)

// JWTAuthMiddleware はJWT認証ミドルウェアを提供します
func JWTAuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorizationヘッダーの取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "認証ヘッダーがありません",
				"code":  model.ErrUnauthorized,
			})
			return
		}

		// Bearer トークンの抽出
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "認証形式が不正です",
				"code":  model.ErrUnauthorized,
			})
			return
		}

		tokenString := parts[1]

		// トークンの検証
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 署名アルゴリズムの確認
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, model.NewUnauthorizedError("不正な署名アルゴリズムです", nil)
			}
			return []byte(cfg.Auth.JWTSecret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "トークンが無効です: " + err.Error(),
				"code":  model.ErrUnauthorized,
			})
			return
		}

		// クレームの取得
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// トークンの有効期限チェック
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "トークンの有効期限が切れています",
						"code":  model.ErrUnauthorized,
					})
					return
				}
			}

			// ユーザーIDをコンテキストに設定
			if userID, ok := claims["user_id"].(float64); ok {
				c.Set("user_id", int(userID))
			}

			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "トークンが無効です",
				"code":  model.ErrUnauthorized,
			})
			return
		}
	}
}

// GenerateToken はJWTトークンを生成します
func GenerateToken(userID int, cfg *config.Config) (string, error) {
	// トークンの有効期限を設定
	expirationTime := time.Now().Add(time.Duration(cfg.Auth.TokenDuration) * time.Minute)

	// クレームを作成
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	// トークンを作成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名
	tokenString, err := token.SignedString([]byte(cfg.Auth.JWTSecret))
	if err != nil {
		return "", model.NewInternalError("トークン生成中にエラーが発生しました", err)
	}

	return tokenString, nil
}
