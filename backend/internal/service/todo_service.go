package service

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
)

type TodoService interface {
	AddNewTodo(ctx context.Context, todoRequest *model.TodoRequest) (*entity.Todo, error)
}
