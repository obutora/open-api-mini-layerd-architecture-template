package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/api/middleware"
	v1 "github.com/obutora/open-api-mini-layerd-architecture-template/internal/api/v1"
	v1handler "github.com/obutora/open-api-mini-layerd-architecture-template/internal/api/v1/handler"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/config"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/repository"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter はAPIルーターを設定します
func SetupRouter(db *sql.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	// リポジトリの初期化
	userRepo := repository.NewUserRepository(db)

	// サービスの初期化
	userService := service.NewUserService(userRepo)
	answerService := service.NewAnswerService()
	paymentsService := service.NewPaymentService()
	questionnaireService := service.NewQuestionnaireService()

	// ミドルウェアの設定
	r.Use(middleware.VersionDetector())

	// 認証が必要なルートのグループ
	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware(cfg))

	// v1 APIルーティング
	v1API := r.Group("/api/v1")
	v1UserHandler := v1handler.NewUserHandler(userService)
	v1AnswerHandler := v1handler.NewAnswerHandler(answerService)
	v1PaymentsHandler := v1handler.NewPaymentHandler(paymentsService)
	v1QuestionnaireHandler := v1handler.NewQuestionnaireHandler(questionnaireService)
	v1.RegisterRoutes(v1API, v1UserHandler, v1AnswerHandler, v1PaymentsHandler, v1QuestionnaireHandler)

	// Swaggerドキュメント
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
