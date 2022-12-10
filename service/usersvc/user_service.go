package usersvc

import (
	"context"
	"errors"

	"github.com/platformsh/template-golang/domain"
	"github.com/platformsh/template-golang/repository"
	"github.com/platformsh/template-golang/service"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) service.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (instance userService) CreateNewUser(ctx context.Context, in domain.CreteUserRequest) error {
	// get user by username
	userCheck, err := instance.userRepo.GetOneByUserName(in.GetUserName())
	if err != nil {
		return err
	}

	if !userCheck.IsEmpty() {
		return errors.New("username sudah digunakan")
	}

	hashedPassword, err := in.GeneratePassword()
	if err != nil {
		return errors.New("failed to generate password")
	}

	// set password
	user := in.ToUser(hashedPassword)

	// save user
	if _, err := instance.userRepo.Create(user); err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (instance userService) Login(ctx context.Context, in domain.LoginRequest) (sessionToken string, err error) {
	// get user by username
	user, err := instance.userRepo.GetOneByUserName(in.GetUserName())
	if err != nil {
		return "", err
	}

	if user.IsEmpty() {
		return "", errors.New("invalid username")
	}

	if err := user.ComparePassword(in.GetPassword()); err != nil {
		return "", errors.New("invalid password")
	}

	// generate sessionToken
	sessionToken = user.GenerateSession()

	// set login at and session token
	user.SetLogin()
	user.SetSessionToken(sessionToken)

	// update user
	if _, err := instance.userRepo.Update(user); err != nil {
		return "", errors.New("failed to update user")
	}

	return sessionToken, nil
}

func (instance userService) Logout(ctx context.Context, in domain.LogoutRequest) error {
	// get user by token
	user, err := instance.userRepo.GetOneBySessionToken(in.GetSessionToken())
	if err != nil {
		return err
	}

	if user.IsEmpty() {
		return errors.New("invalid session token")
	}

	user.SetSessionTokenToBeEmpty()
	user.SetLogout()
	if _, err := instance.userRepo.Update(user); err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

func (instance userService) CheckToken(ctx context.Context, token string) (bool, error) {
	// get user by token
	user, err := instance.userRepo.GetOneBySessionToken(token)
	if err != nil {
		return false, err
	}

	return !user.IsEmpty(), nil
}
