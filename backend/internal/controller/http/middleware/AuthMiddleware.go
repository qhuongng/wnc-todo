package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/http_common"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/utils/authentication"
	"net/http"
)

type tokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func VerifyTokenMiddleware(c *gin.Context) {
	//accessToken, err := c.Request.Cookie("access_token")
	accessToken := c.GetHeader("access_token")
	//var body *tokenRequest
	//if err := c.ShouldBindJSON(&body); err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, http_common.NewErrorResponse(http_common.Error{
	//		Message: "Missing token", Field: "token", Code: http_common.ErrorResponseCode.InvalidRequest,
	//	}))
	//	return
	//}
	//accessToken := body.AccessToken
	//refreshToken := body.RefreshToken

	//if !accessExists || !refreshExists {
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, http_common.NewErrorResponse(http_common.Error{
	//		Message: "Missing token", Field: "token", Code: http_common.ErrorResponseCode.Unauthorized,
	//	}))
	//	return
	//}
	//check accesstoken
	userId, err := authentication.VerifyToken(accessToken, "access")
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			refreshToken := c.GetHeader("refresh_token")
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
