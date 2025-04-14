package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/api/v1/handler"
)

// RegisterRoutes はv1 APIのルートを登録します
func RegisterRoutes(r *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) {
	// ユーザー関連エンドポイント
	users := r.Group("/users")
	{
		users.GET("", userHandler.ListUsers)         // GET /api/v1/users
		users.POST("", userHandler.CreateUser)       // POST /api/v1/users
		users.GET("/:id", userHandler.GetUser)       // GET /api/v1/users/:id
		users.PUT("/:id", userHandler.UpdateUser)    // PUT /api/v1/users/:id
		users.DELETE("/:id", userHandler.DeleteUser) // DELETE /api/v1/users/:id
	}

	// 製品関連エンドポイント
	products := r.Group("/products")
	{
		products.GET("", productHandler.ListProducts)         // GET /api/v1/products
		products.POST("", productHandler.CreateProduct)       // POST /api/v1/products
		products.GET("/:id", productHandler.GetProduct)       // GET /api/v1/products/:id
		products.PUT("/:id", productHandler.UpdateProduct)    // PUT /api/v1/products/:id
		products.DELETE("/:id", productHandler.DeleteProduct) // DELETE /api/v1/products/:id
	}
}
