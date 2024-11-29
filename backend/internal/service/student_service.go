package service

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
)

type StudentService interface {
	GetAllStudent(ctx context.Context) []model.Student
}
