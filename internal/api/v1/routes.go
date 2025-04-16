package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/api/middleware"
	"github.com/obutora/open-api-mini-layerd-architecture-template/internal/api/v1/handler"
)

// RegisterRoutes はv1 APIのルートを登録します
func RegisterRoutes(r *gin.RouterGroup, userHandler *handler.UserHandler, answerHandler *handler.AnswerHandler, paymentsHandler *handler.PaymentHandler, questionnaireHandler *handler.QuestionnaireHandler) {
	// ユーザー関連エンドポイント
	users := r.Group("/users")
	{
		users.GET("", userHandler.ListUsers)         // GET /api/v1/users
		users.POST("", userHandler.CreateUser)       // POST /api/v1/users
		users.GET("/:id", userHandler.GetUser)       // GET /api/v1/users/:id
		users.PUT("/:id", userHandler.UpdateUser)    // PUT /api/v1/users/:id
		users.DELETE("/:id", userHandler.DeleteUser) // DELETE /api/v1/users/:id
	}

	questionnaires := r.Group("/questionnaires")
	{
		questionnaires.GET("/list", middleware.SimpleApiKeyAuthMiddleware(), questionnaireHandler.ListQuestionnaires)           // GET /api/v1/questionnaires/list
		questionnaires.GET("/:id/status", middleware.SimpleApiKeyAuthMiddleware(), questionnaireHandler.GetQuestionnaireStatus) // GET /api/v1/questionnaires/{id}/status [get]

		questionnaires.GET("/:id", questionnaireHandler.GetQuestionnaire)                                               // GET /api/v1/questionnaires/:id?hash=xxx
		questionnaires.PATCH("/:id", middleware.SimpleApiKeyAuthMiddleware(), questionnaireHandler.UpdateQuestionnaire) // PATCH /api/v1/questionnaires
	}

	answers := r.Group("/answer")
	{
		answers.GET("/completion", answerHandler.Completion) // GET /api/v1/answer/completion
	}

	// TODO: 別コンテナにするかも
	payments := r.Group("/payments")
	{
		payments.Use(middleware.SimpleApiKeyAuthMiddleware())
		payments.GET("", paymentsHandler.GetPayments) // GET /api/v1/payments
	}

}
