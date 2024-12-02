package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/http_common"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/utils/authentication"
	"net/http"
)

func VerifyTokenMiddleware(c *gin.Context) {
	accessToken, err := c.Request.Cookie("access_token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "access_token", Code: http_common.ErrorResponseCode.Unauthorized,
		}))
		return
	}
	//check accesstoken
	userId, err := authentication.VerifyToken(accessToken.Value, "access")
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			//check refreshtoken
			refreshToken, err := c.Request.Cookie("refresh_token")
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
					Message: err.Error(), Field: "refresh_token", Code: http_common.ErrorResponseCode.Unauthorized,
				}))
				return
			}
			userId, err = authentication.VerifyToken(refreshToken.Value, "refresh")
			if err != nil {
				if errors.Is(err, jwt.ErrTokenExpired) {
					c.AbortWithStatusJSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
						Message: "Token expired", Field: "refresh_token", Code: http_common.ErrorResponseCode.Unauthorized,
					}))
					return
				}
				c.AbortWithStatusJSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
					Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.Unauthorized,
				}))
				return
			}
			c.Set("userId", userId)
			c.Set("resetAccessToken", true)
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
			Message: err.Error(), Field: "", Code: http_common.ErrorResponseCode.Unauthorized,
		}))
		return
	}
	c.Set("userId", userId)
	c.Set("resetAccessToken", false)
}
