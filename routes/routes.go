package routes

import (
	"main/controller"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	// ref := "https://github.dev/cjaewon/echo-gorm-example"
	e := echo.New()

	// implement middleware
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 30 * time.Second,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))

	userRoutes := e.Group("/users")
	userRoutes.GET("/:id", controller.GetUserController)
	userRoutes.POST("/", controller.PostUserController)
	userRoutes.PUT("/:id", controller.UpdateUserController)

	quizRoutes := e.Group("/quizzes")
	quizRoutes.GET("/", controller.GetQuizzesController)
	quizRoutes.POST("/", controller.CreateQuizController)
	quizRoutes.PUT("/:id", controller.UpdateQuizController)
	quizRoutes.GET("/:id", controller.GetQuizController)
	quizRoutes.DELETE("/:id", controller.DeleteQuizController)

	questionRoutes := e.Group("/quizzes/:quizId/questions")
	questionRoutes.GET("/", controller.GetAllQuestionController)
	questionRoutes.GET("/:questionId", controller.GetQuestionController)
	questionRoutes.POST("/", controller.CreateQuestionController)
	questionRoutes.PUT("/:questionId", controller.UpdateQuestionController)
	questionRoutes.DELETE("/:questionId", controller.DeleteQuestionController)

	quizResultRoutes := e.Group("/users/:user_id/quizzes/:quiz_id/result")
	quizResultRoutes.GET("/", controller.GetQuizResultController)
	quizResultRoutes.POST("/", controller.AddQuizResultController)
	return e
}
