package serviceimplement

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/repository"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/service"
)

type TodoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) service.TodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (service *TodoService) AddNewTodo(ctx context.Context, todoRequest *model.TodoRequest) (*entity.Todo, error) {
	todo := &entity.Todo{
		Content:   todoRequest.Content,
		UserId:    todoRequest.UserId,
		Completed: false,
	}
	err := service.todoRepository.AddNewTodo(ctx, todo)
	if err != nil {
		return nil, err
	}
	newTodo, err := service.todoRepository.GetListTodo(ctx, todoRequest.UserId, todoRequest.Content)
	if err != nil {
		return nil, err
	}
	return &newTodo[0], nil
}
