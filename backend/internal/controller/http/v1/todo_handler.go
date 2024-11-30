package v1

import "github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/service"

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}
