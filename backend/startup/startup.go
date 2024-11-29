package startup

import (
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/controller"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/database"
)

func registerDependencies() *controller.ApiContainer {
	// Open database connection
	db := database.Open()

	return internal.InitializeContainer(db)
}

func Execute() {
	container := registerDependencies()
	container.HttpServer.Run()
}
