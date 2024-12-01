package repositoryimplement

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/database"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db database.Db) repository.UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) AddNewUser(c context.Context, user *entity.User) error {
	err := repo.db.WithContext(c).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserRepository) GetUserByName(c context.Context, username string) (*entity.User, error) {
	var user entity.User
	err := repo.db.WithContext(c).First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *UserRepository) GetUserById(c context.Context, userId int64) (*entity.User, error) {
	var user entity.User
	err := repo.db.WithContext(c).First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *UserRepository) UpdateUser(c context.Context, user *entity.User, userId int64) (*entity.User, error) {
	if err := repo.db.WithContext(c).Model(&entity.User{Id: userId}).Updates(&user).Error; err != nil {
		return nil, err
	}
	var updateUser entity.User
	err := repo.db.WithContext(c).First(&updateUser, userId).Error
	if err != nil {
		return nil, err
	}
	return &updateUser, nil
}
