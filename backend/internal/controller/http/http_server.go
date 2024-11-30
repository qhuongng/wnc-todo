package http

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	v1 "github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/controller/http/v1"
)

type Server struct {
	userHandler *v1.UserHandler
	todoHandler *v1.TodoHandler
}

func NewServer(userHandler *v1.UserHandler, todoHandler *v1.TodoHandler) *Server {
	return &Server{userHandler: userHandler, todoHandler: todoHandler}
}

func (s *Server) Run() {
	router := gin.New()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	httpServerInstance := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	v1.MapRoutes(router, s.userHandler, s.todoHandler)
	err := httpServerInstance.ListenAndServe()
	if err != nil {
		return
	}
	fmt.Println("Server running at " + httpServerInstance.Addr)
}
