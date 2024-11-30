package repositoryimplement

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/database"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/repository"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db database.Db) repository.TodoRepository {
	return &TodoRepository{db: db}
}

func (repo *TodoRepository) AddNewTodo(c context.Context, todo *entity.Todo) error {
	return nil
}
