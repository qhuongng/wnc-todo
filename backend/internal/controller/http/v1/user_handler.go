package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/http_common"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/service"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

type LoginResponse struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// @Summary Login
// @Description Login to account
// @Tags Users
// @Accept json
// @Param request body model.UserRequest true "User payload"
// @Produce  json
// @Router /users/login [post]
// @Success 200 {object} http_common.HttpResponse[string]
// @Failure 400 {object} http_common.HttpResponse[any]
// @Failure 500 {object} http_common.HttpResponse[any]
func (handler *UserHandler) Login(c *gin.Context) {
	var req *model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	user, err := handler.userService.Login(c, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.Unauthorized,
		}))
		return
	}
	//create accesstoken
	accessToken, err := handler.userService.CreateToken(c, user.Id, "access")
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	// set access token
	c.Header("access_token", accessToken)
	// set refresh token
	c.Header("refresh_token", user.RefeshToken)
	//c.SetCookie("access_token", accessToken, constants.COOKIE_DURATION, "/", "", false, true)
	//// set refresh token
	//c.SetCookie("refresh_token", user.RefeshToken, constants.COOKIE_DURATION, "/", "", false, true)

	c.JSON(http.StatusOK, http_common.NewSuccessResponse[LoginResponse](&LoginResponse{Username: user.Username, AccessToken: accessToken, RefreshToken: user.RefeshToken}))
}

// @Summary Register
// @Description Register to account
// @Tags Users
// @Accept json
// @Param request body model.UserRequest true "User payload"
// @Produce  json
// @Router /users/register [post]
// @Success 200 {object} http_common.HttpResponse[string]
// @Failure 400 {object} http_common.HttpResponse[any]
// @Failure 500 {object} http_common.HttpResponse[any]
func (handler *UserHandler) Register(c *gin.Context) {
	var req *model.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InvalidRequest,
		}))
		return
	}
	user, err := handler.userService.Register(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.Unauthorized,
		}))
		return
	}
	accessToken, err := handler.userService.CreateToken(c, user.Id, "access")
	if err != nil {
		c.JSON(http.StatusInternalServerError, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.InternalServerError,
		}))
		return
	}
	// set access token
	c.Header("access_token", accessToken)
	// set refresh token
	c.Header("refresh_token", user.RefeshToken)
	//c.SetCookie("access_token", accessToken, constants.COOKIE_DURATION, "/", "", false, true)
	//// set refresh token
	//c.SetCookie("refresh_token", user.RefeshToken, constants.COOKIE_DURATION, "/", "", false, true)
	c.JSON(http.StatusOK, http_common.NewSuccessResponse[LoginResponse](&LoginResponse{Username: user.Username, AccessToken: accessToken, RefreshToken: user.RefeshToken}))
}

func (handler *UserHandler) Logout(c *gin.Context) {

}
