package service

import (
	"context"
	"errors"

	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/adapter/config"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/core/port"
	"github.com/yehezkiel1086/go-rabbitmq-email-notification/user-service/internal/core/util"
)

type AuthService struct {
	conf *config.JWT
	userRepo port.UserRepository
}

func NewAuthService(conf *config.JWT, userRepo port.UserRepository) *AuthService {
	return &AuthService{
		conf,
		userRepo,
	}
}

func (as *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	// check email
	user, err := as.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	// check is_verified
	if !user.IsVerified {
		return "", errors.New("user is not verified")
	}

	// check password
	if err := util.CompareHashedPwd(user.Password, password); err != nil {
		return "", err
	}

	// generate jwt token
	token, err := util.GenerateJWTToken(as.conf, user)
	if err != nil {
		return "", err
	}

	return token, nil
}
