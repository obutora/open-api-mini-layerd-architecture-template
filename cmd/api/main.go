package main

// @title API Documentation
// @version 1.0
// @description API documentation for the application
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Bearer Tokenによる認証。'Bearer 'プレフィックスの後にJWTトークンを入力してください。

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/obutora/open-api-mini-layerd-architecture-template/docs"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/api"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/config"
)

func main() {
	// Swaggerドキュメントの初期化
	docs.SwaggerInfo.Title = "API Documentation"
	docs.SwaggerInfo.Description = "API documentation for the application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	// 設定の読み込み
	cfg := config.Load()

	// データベース接続
	// db, err := sql.Open(cfg.Database.Driver, cfg.Database.DSN())
	// if err != nil {
	// 	log.Fatalf("データベース接続エラー: %v", err)
	// }
	// defer db.Close()
	// スキーママイグレーション

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 接続テスト
	if err := db.Ping(); err != nil {
		log.Fatalf("データベース接続テストエラー: %v", err)
	}

	// ルーターの設定
	router := api.SetupRouter(db, cfg)

	// サーバー起動
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("サーバーを起動しています: %s", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: router,
		// ReadTimeout:  cfg.Server.ReadTimeout,
		// WriteTimeout: cfg.Server.WriteTimeout,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("サーバー起動エラー: %v", err)
	}
}
