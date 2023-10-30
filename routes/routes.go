package routes

import (
	"main/controller"
	"os"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
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

	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET_KEY"))))

	e.POST("/users/register", controller.RegisterController)
	e.POST("/users/login", controller.LoginController)
	eAuth.PUT("/users/update/:id", controller.UpdateUserController)

	eAuth.GET("/quizzes", controller.GetQuizzesController)
	eAuth.POST("/quizzes", controller.CreateQuizController)
	eAuth.PUT("/quizzes/:id", controller.UpdateQuizController)
	eAuth.GET("/quizzes/:id", controller.GetQuizController)
	eAuth.DELETE("/quizzes/:id", controller.DeleteQuizController)

	eAuth.GET("/quizzes/:quizId/questions", controller.GetAllQuestionController)
	eAuth.GET("/quizzes/:quizId/questions/:questionId", controller.GetQuestionController)
	eAuth.POST("/quizzes/:quizId/questions", controller.CreateQuestionController)
	eAuth.PUT("/quizzes/:quizId/questions/:questionId", controller.UpdateQuestionController)
	eAuth.DELETE("/quizzes/:quizId/questions/:questionId", controller.DeleteQuestionController)

	eAuth.GET("/quizzes/:quiz_id/result", controller.GetQuizResultController)
	eAuth.POST("/quizzes/:quiz_id/result", controller.AddQuizResultController)
	return e
}
