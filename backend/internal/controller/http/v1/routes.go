package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/controller/http/middleware"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, userHandler *UserHandler, todoHandler *TodoHandler) {
	router.Use(middleware.CorsMiddleware())
	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.POST("/login", userHandler.Login)
		}
		todos := v1.Group("/todos")
		todos.Use(middleware.VerifyTokenMiddleware)
		{
			todos.POST("/", todoHandler.Add)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
