package middleware

import (
	"github.com/gin-gonic/gin"
)

// APIVersion はAPIバージョンを表します
type APIVersion string

const (
	// V1 はAPIバージョン1を表します
	V1 APIVersion = "v1"
	// V2 はAPIバージョン2を表します
	V2 APIVersion = "v2"
)

// VersionDetector はリクエストからAPIバージョンを検出するミドルウェアです
func VersionDetector() gin.HandlerFunc {
	return func(c *gin.Context) {
		// パスからバージョンを抽出
		// path := c.Request.URL.Path
		version := V1 // デフォルトはv1

		// if strings.Contains(path, "/api/v2/") {
		// 	version = V2
		// }

		// バージョン情報をコンテキストに設定
		c.Set("api_version", version)
		c.Next()
	}
}

// GetRequestedVersion はリクエストされたAPIバージョンを取得します
func GetRequestedVersion(c *gin.Context) APIVersion {
	version, exists := c.Get("api_version")
	if !exists {
		return V1 // デフォルトはv1
	}
	return version.(APIVersion)
}
