package authentication

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

var secretKeyAccess = os.Getenv("JWT_ACCESS_SECRET_KEY")
var secretKeyRefresh = os.Getenv("JWT_REFRESH_SECRET_KEY")

func GenerateToken(userId int64, expTime time.Time, tokenType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId": userId,
			"exp":    expTime.Unix(),
		})
	var key string
	if tokenType == "access" {
		key = secretKeyAccess
	} else if tokenType == "refresh" {
		key = secretKeyRefresh
	} else {
		return "", errors.New("Invalid token type")
	}
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func VerifyToken(tokenString string, tokenType string) (string, error) {
	var key string

	if tokenType == "access" {
		key = secretKeyAccess
	} else if tokenType == "refresh" {
		key = secretKeyRefresh
	} else {
		return "", errors.New("Invalid token type")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		if tokenType == "access" {
			return "", fmt.Errorf("invalid accessToken")
		} else {
			return "", fmt.Errorf("invalid refreshToken")
		}
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userId, ok := claims["userId"].(string); ok {
			return userId, nil
		}
		return "", errors.New("userId not found in token")
	}

	return "", errors.New("failed to parse token claims")
}
