package repository

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
)

type UserRepository interface {
	AddNewUser(c context.Context, user *entity.User) error
	GetUserByName(c context.Context, username string) (*entity.User, error)
	GetUserById(c context.Context, userId string) (*entity.User, error)
	UpdateUser(c context.Context, user *entity.User, userId int64) (*entity.User, error)
}
