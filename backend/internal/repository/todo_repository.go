package repository

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
)

type TodoRepository interface {
	AddNewTodo(c context.Context, todo *entity.Todo) (int64, error)
	UpdateTodo(c context.Context, todo *entity.Todo, todoId int64, userId int64) (*entity.Todo, error)
	GetListTodo(c context.Context, userId int64, searchText string) ([]entity.Todo, error)
}
