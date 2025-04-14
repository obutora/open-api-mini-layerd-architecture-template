
## ディレクトリ構造
簡易的かつシンプルなレイヤードアーキテクチャを採用しています。
依存先が比較的明瞭であることに加えて、ラーニングコストが低く、平易で、LLMにも読みやすい構造を意識しています。

```
project-root/
│
├── cmd/
│   └── api/
│       └── main.go                  # エントリーポイント
│
├── internal/
│   ├── model/                       # ドメインモデル
│   │   ├── user.go                  # ユーザーモデル
│   │   ├── product.go               # 製品モデル
│   │   └── error.go                 # エラー定義
│   │
│   ├── repository/                  # リポジトリインターフェースと実装
│   │   ├── user_repository.go       # ユーザーリポジトリ
│   │   └── product_repository.go    # 製品リポジトリイン
│   │
│   ├── service/                     # ビジネスロジック
│   │   ├── user_service.go          # ユーザーサービス
│   │   └── product_service.go       # 製品サービス
│   │
│   ├── config/                      # 設定
│   │   └── config.go                # アプリケーション設定
│   │
│   ├── api/                         # API関連
│   │   ├── router.go                # メインルーター定義
│   │   │
│   │   ├── middleware/
│   │   │   ├── auth.go              # 認証ミドルウェア
│   │   │   └── version.go           # バージョン識別ミドルウェア
│   │   │
│   │   ├── v1/                      # v1 API実装
│   │   │   ├── handler/
│   │   │   │   ├── user_handler.go  # ユーザーハンドラー
│   │   │   │   └── product_handler.go # 製品ハンドラー
│   │   │   │
│   │   │   ├── dto/                 # DTOモデル
│   │   │   │   ├── user.go          # リクエスト/レスポンスDTO
│   │   │   │   └── product.go       # リクエスト/レスポンスDTO
│   │   │   │
│   │   │   └── routes.go            # v1ルート定義
│   │   │
│   │   └── v2/                      # v2 API実装
│   │       ├── handler/
│   │       │   ├── user_handler.go
│   │       │   └── product_handler.go
│   │       │
│   │       ├── dto/
│   │       │   ├── user.go
│   │       │   └── product.go
│   │       │
│   │       └── routes.go            # v2ルート定義
│   │
│   └── util/                        # ユーティリティ
│       ├── logger.go
│       └── validator.go
│
├── docs/                            # Swagger生成ドキュメント
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
│
├── i18n/                            # 多言語対応リソース
│   ├── en/
│   │   └── messages.json
│   └── ja/
│       └── messages.json
│
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 主要なコンポーネントの役割

### 1. Model層
ビジネスエンティティとドメインロジックを定義します。

```go
// internal/model/user.go
package model

type User struct {
    ID       int
    Name     string
    Email    string
    // 他のフィールド
}

// ドメインロジックを持つメソッドを追加可能
func (u *User) IsValid() bool {
    return u.Email != "" && u.Name != ""
}
```

### 2. Repository層
データアクセスのインターフェースと実装を提供します。

```go
// internal/repository/user_repository.go
package repository

import "myapp/internal/model"

type UserRepository interface {
    GetByID(id int) (*model.User, error)
    Create(user *model.User) error
    // 他のメソッド
}

type UserRepositoryImpl struct {
    db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
    return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) GetByID(id int) (*model.User, error) {
    // 実装
}
```

### 3. Service層
ビジネスロジックを実装します。複数のリポジトリを使用する処理などを担当。

```go
// internal/service/user_service.go
package service

import "myapp/internal/model"
import "myapp/internal/repository"

type UserService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) GetUser(id int) (*model.User, error) {
    return s.userRepo.GetByID(id)
}
```

### 4. API/Handler層
HTTPリクエストの処理とレスポンスの返却を担当します。バージョン管理はここで行います。

```go
// internal/api/v1/handler/user_handler.go
package handler

import (
    "myapp/internal/model"
    "myapp/internal/service"
    "myapp/internal/api/v1/dto"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

// @Summary ユーザー取得
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    
    user, err := h.userService.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
        return
    }
    
    // モデルをDTOに変換
    response := dto.UserResponse{
        ID:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }
    
    c.JSON(http.StatusOK, response)
}
```

### v2のハンドラー例（拡張機能）

```go
// internal/api/v2/handler/user_handler.go
package handler

import (
    "myapp/internal/service"
    "myapp/internal/api/v2/dto"
)

type UserHandler struct {
    userService *service.UserService
}

// @Router /api/v2/users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    
    user, err := h.userService.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません", "code": "NOT_FOUND"})
        return
    }
    
    // v2拡張DTOに変換
    response := dto.UserResponse{
        ID:        user.ID,
        Name:      user.Name,
        Email:     user.Email,
        CreatedAt: user.CreatedAt,  // v2で追加されたフィールド
    }
    
    c.JSON(http.StatusOK, response)
}
```

### ルーティング設定

```go
// internal/api/router.go
func SetupRouter(db *sql.DB) *gin.Engine {
    r := gin.Default()
    
    // リポジトリ
    userRepo := mysql.NewUserRepository(db)
    
    // サービス
    userService := service.NewUserService(userRepo)
    
    // バージョンミドルウェア
    r.Use(middleware.VersionDetector())
    
    // v1ルーティング
    v1API := r.Group("/api/v1")
    v1UserHandler := v1handler.NewUserHandler(userService)
    v1.RegisterRoutes(v1API, v1UserHandler)
    
    // v2ルーティング
    v2API := r.Group("/api/v2")
    v2UserHandler := v2handler.NewUserHandler(userService)
    v2.RegisterRoutes(v2API, v2UserHandler)
    
    return r
}
```

```go
// internal/api/v1/routes.go
package v1

import (
	"myapp/internal/api/v1/handler"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes はv1 APIのルートを登録します
func RegisterRoutes(r *gin.RouterGroup, userHandler *handler.UserHandler, productHandler *handler.ProductHandler) {
	// ユーザー関連エンドポイント
	users := r.Group("/users")
	{
		users.GET("", userHandler.ListUsers)        // GET /api/v1/users
		users.POST("", userHandler.CreateUser)      // POST /api/v1/users
		users.GET("/:id", userHandler.GetUser)      // GET /api/v1/users/:id
		users.PUT("/:id", userHandler.UpdateUser)   // PUT /api/v1/users/:id
		users.DELETE("/:id", userHandler.DeleteUser) // DELETE /api/v1/users/:id
	}

	// 製品関連エンドポイント
	products := r.Group("/products")
	{
		products.GET("", productHandler.ListProducts)        // GET /api/v1/products
		products.POST("", productHandler.CreateProduct)      // POST /api/v1/products
		products.GET("/:id", productHandler.GetProduct)      // GET /api/v1/products/:id
		products.PUT("/:id", productHandler.UpdateProduct)   // PUT /api/v1/products/:id
		products.DELETE("/:id", productHandler.DeleteProduct) // DELETE /api/v1/products/:id
	}
}
```


## 自動生成関連

see also [generate.go](./internal/generate.go)
```
go generate .
```
see also [Makefile](./Makefile)
```
make docs
```

