package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/controller/http/middleware"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func MapRoutes(router *gin.Engine, userHandler *UserHandler, todoHandler *TodoHandler) {
	router.Use(middleware.CorsMiddleware())
	//router.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	//	AllowHeaders:     []string{"Content-Type", "Authorization", "access-token", "refresh-token"},
	//	AllowCredentials: true,
	//	MaxAge:           12 * time.Hour,
	//}))
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
			todos.PUT("/:id", todoHandler.Update)
			todos.GET("/", todoHandler.GetList)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
