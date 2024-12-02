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
	//accessToken, err := c.Request.Cookie("access_token")
	accessToken := c.GetHeader("access_token")
	if accessToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
			Message: "Missing access token", Field: "access_token", Code: http_common.ErrorResponseCode.Unauthorized,
		}))
		return
	}
	//check accesstoken
	userId, err := authentication.VerifyToken(accessToken, "access")
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			//check refreshtoken
			//refreshToken, err := c.Request.Cookie("refresh_token")
			refreshToken := c.GetHeader("refresh_token")
			if refreshToken == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
					Message: "Missing refresh token", Field: "refresh_token", Code: http_common.ErrorResponseCode.Unauthorized,
				}))
				return
			}
			userId, err = authentication.VerifyToken(refreshToken, "refresh")
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
