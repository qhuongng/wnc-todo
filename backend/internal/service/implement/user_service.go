package serviceimplement

import (
	"context"
	"errors"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/entity"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/domain/model"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/repository"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/internal/service"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/utils/authentication"
	"github.com/qhuongng/wnc-todo/tree/add-redux/backend/utils/constants"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &UserService{userRepository: userRepository}
}
func (service *UserService) Login(ctx context.Context, userRequest *model.UserRequest) (*entity.User, error) {
	user, err := service.userRepository.GetUserByName(ctx, userRequest.Username)
	if err != nil {
		return nil, err
	}
	//check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) Register(ctx context.Context, userRequest *model.UserRequest) (*entity.User, error) {
	//check invalid
	user, err := service.userRepository.GetUserByName(ctx, userRequest.Username)
	if user != nil {
		return nil, errors.New("user already exists")
	}
	//hash password before save
	hashPW, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user = &entity.User{
		Username: userRequest.Username,
		Password: string(hashPW),
	}
	err = service.userRepository.AddNewUser(ctx, user)
	if err != nil {
		return nil, err
	}
	//setup refreshtoken
	currentUser, err := service.userRepository.GetUserByName(ctx, userRequest.Username)
	if err != nil {
		return nil, err
	}
	_, err = service.CreateToken(ctx, currentUser.Id, "refresh")
	if err != nil {
		return nil, err
	}
	return currentUser, nil
}

func (service *UserService) CreateToken(ctx context.Context, userId int64, tokenType string) (string, error) {
	currentUser, err := service.userRepository.GetUserById(ctx, userId)
	if err != nil {
		return "", err
	}
	if tokenType == "refresh" {
		refreshTokenTime := time.Now().Add(constants.REFRESH_TOKEN_DURATION)
		refreshToken, err := authentication.GenerateToken(currentUser, refreshTokenTime, "refresh")
		if err != nil {
			return "", err
		}
		currentUser.RefeshToken = refreshToken
		//update user
		_, err = service.userRepository.UpdateUser(ctx, currentUser, currentUser.Id)
		if err != nil {
			return "", err
		}
		return refreshToken, nil
	} else if tokenType == "access" {
		accessTokenTime := time.Now().Add(constants.ACCESS_TOKEN_DURATION)
		accessToken, err := authentication.GenerateToken(currentUser, accessTokenTime, "access")
		if err != nil {
			return "", err
		}
		return accessToken, nil
	} else {
		return "", errors.New("invalid token type")
	}
}
