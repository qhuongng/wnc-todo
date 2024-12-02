package service

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
)

type TodoService interface {
	AddNewTodo(ctx context.Context, todoRequest *model.TodoRequest) (*entity.Todo, error)
	UpdateTodo(ctx context.Context, todoRequest *model.TodoRequest, todoId int64) (*entity.Todo, error)
	GetListTodo(ctx context.Context, userId int64, searchText string) ([]entity.Todo, error)
}
