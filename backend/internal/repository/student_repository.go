package repository

import (
	"context"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
)

type StudentRepository interface {
	GetAllStudentQuery(c context.Context) []entity.Student
}
