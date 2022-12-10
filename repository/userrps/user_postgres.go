package userrps

import (
	"errors"

	"github.com/platformsh/template-golang/domain"
	"github.com/platformsh/template-golang/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	maria *gorm.DB
}

func NewUserRepository(maria *gorm.DB) repository.UserRepository {
	return &userRepo{
		maria: maria,
	}
}

func (instance userRepo) Create(user *domain.User) (*domain.User, error) {
	if err := instance.maria.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (instance userRepo) Update(user *domain.User) (*domain.User, error) {
	if err := instance.maria.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (instance userRepo) GetOneByUserName(userName string) (*domain.User, error) {
	var user *domain.User

	if err := instance.maria.Where("user_name = ?", userName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func (instance userRepo) GetOneBySessionToken(sessiontToken string) (*domain.User, error) {
	var user *domain.User

	if err := instance.maria.Where("session_token = ?", sessiontToken).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
