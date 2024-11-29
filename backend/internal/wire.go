//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/controller"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/controller/http"
	v1 "github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/controller/http/v1"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/database"
	repositoryimplement "github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/repository/implement"
	serviceimplement "github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/service/implement"
)

var container = wire.NewSet(
	controller.NewApiContainer,
)

// may have grpc server in the future
var serverSet = wire.NewSet(
	http.NewServer,
)

// handler === controller | with service and repository layers to form 3 layers architecture
var handlerSet = wire.NewSet(
	v1.NewStudentHandler,
)

var serviceSet = wire.NewSet(
	serviceimplement.NewStudentService,
)

var repositorySet = wire.NewSet(
	repositoryimplement.NewStudentRepository,
)

func InitializeContainer(
	db database.Db,
) *controller.ApiContainer {
	wire.Build(serverSet, handlerSet, serviceSet, repositorySet, container)
	return &controller.ApiContainer{}
}
