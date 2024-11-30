package repository

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
)

type TodoRepository interface {
	AddNewTodo(c context.Context, todo *entity.Todo) error
}
