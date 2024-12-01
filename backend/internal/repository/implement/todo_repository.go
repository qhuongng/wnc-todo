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
	err := repo.db.WithContext(c).Create(todo).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *TodoRepository) UpdateTodo(c context.Context, todo *entity.Todo, todoId int64, userId int64) (*entity.Todo, error) {
	if err := repo.db.WithContext(c).Model(&entity.Todo{Id: todoId, UserId: userId}).Updates(&todo).Error; err != nil {
		return nil, err
	}
	var updatedTodo entity.Todo
	err := repo.db.WithContext(c).First(&updatedTodo, todoId).Error
	if err != nil {
		return nil, err
	}
	return &updatedTodo, nil
}

func (repo *TodoRepository) GetListTodo(c context.Context, userId int64, searchText string) ([]entity.Todo, error) {
	var todoList []entity.Todo

	query := repo.db.WithContext(c).Where("user_id = ?", userId)
	if searchText != "" {
		query = query.Where("content LIKE ?", "%"+searchText+"%")
	}

	if err := query.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}
