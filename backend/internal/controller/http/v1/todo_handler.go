package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/http_common"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/service"
	"net/http"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (handler *TodoHandler) Add(c *gin.Context) {
	var req *model.TodoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
			Message: "Invalid token", Field: "userId", Code: http_common.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	req.UserId = userId.(int64)
	newTodo, err := handler.todoService.AddNewTodo(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
	}
	c.JSON(http.StatusOK, http_common.NewSuccessResponse[*entity.Todo](&newTodo))
}

func (handler *TodoHandler) Update(c *gin.Context) {

}

func (handler *TodoHandler) GetList(c *gin.Context) {

}
