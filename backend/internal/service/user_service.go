package service

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
)

type UserService interface {
	Login(ctx context.Context, userRequest *model.UserRequest) (*entity.User, error)
	Register(ctx context.Context, userRequest *model.UserRequest) (*entity.User, error)
	CreateToken(ctx context.Context, userId int64, tokenType string) (string, error)
}
