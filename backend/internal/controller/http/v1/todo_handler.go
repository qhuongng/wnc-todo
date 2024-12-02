package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/http_common"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/service"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/utils/constants"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	todoService service.TodoService
	userService service.UserService
}

func NewTodoHandler(todoService service.TodoService, userService service.UserService) *TodoHandler {
	return &TodoHandler{todoService: todoService, userService: userService}
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
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: "Invalid token", Field: "userId", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	//check token
	isExpiredToken, exists := c.Get("resetAccessToken")
	if !exists {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: "Invalid token", Field: "Regenerate token", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	if isExpiredToken == true {
		newToken, err := handler.userService.CreateToken(c, userId.(int64), "access")
		if err != nil {
			c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
				Message: err.Error(), Field: "Regenerate token", Code: http_common.ErrorResponseCode.InternalServerError,
			}))
			return
		}
		c.SetCookie("access_token", newToken, constants.COOKIE_DURATION, "/", "", false, true)
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
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(http_common.Error{
			Message: "Missing param id", Field: "id", Code: http_common.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	todoId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "id", Code: http_common.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	var req *model.TodoRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: "Invalid token", Field: "userId", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	//check token
	isExpiredToken, exists := c.Get("resetAccessToken")
	if !exists {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: "Invalid token", Field: "Regenerate token", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	if isExpiredToken == true {
		newToken, err := handler.userService.CreateToken(c, userId.(int64), "access")
		if err != nil {
			c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
				Message: err.Error(), Field: "Regenerate token", Code: http_common.ErrorResponseCode.InternalServerError,
			}))
			return
		}
		c.SetCookie("access_token", newToken, constants.COOKIE_DURATION, "/", "", false, true)
	}
	req.UserId = userId.(int64)
	updatedUser, err := handler.todoService.UpdateTodo(c, req, todoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	c.JSON(http.StatusOK, http_common.NewSuccessResponse[*entity.Todo](&updatedUser))
}

func (handler *TodoHandler) GetList(c *gin.Context) {
	filter := c.Query("filter")
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: "Invalid token", Field: "userId", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	//check token
	isExpiredToken, exists := c.Get("resetAccessToken")
	if !exists {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: "Invalid token", Field: "Regenerate token", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	if isExpiredToken == true {
		newToken, err := handler.userService.CreateToken(c, userId.(int64), "access")
		if err != nil {
			c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
				Message: err.Error(), Field: "Regenerate token", Code: http_common.ErrorResponseCode.InternalServerError,
			}))
			return
		}
		c.SetCookie("access_token", newToken, constants.COOKIE_DURATION, "/", "", false, true)
	}

	todoList, err := handler.todoService.GetListTodo(c, userId.(int64), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
	}
	c.JSON(http.StatusOK, http_common.NewSuccessResponse[[]entity.Todo](&todoList))
}
