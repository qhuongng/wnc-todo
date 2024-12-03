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
		v1.POST("/todos", middleware.VerifyTokenMiddleware, todoHandler.Add)
		v1.PUT("/todos/:id", middleware.VerifyTokenMiddleware, todoHandler.Update)
		v1.GET("/todos", middleware.VerifyTokenMiddleware, todoHandler.GetList)
		v1.POST("/users/register", userHandler.Register)
		v1.POST("/users/login", userHandler.Login)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
