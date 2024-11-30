package serviceimplement

import (
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/repository"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/service"
)

type TodoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) service.TodoService {
	return &TodoService{todoRepository: todoRepository}
}
