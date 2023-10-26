package routes

import (
	"main/controller"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

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
