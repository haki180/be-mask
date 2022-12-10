package repository

import (
	"github.com/platformsh/template-golang/domain"
)

type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	GetOneByUserName(userName string) (*domain.User, error)
	GetOneBySessionToken(sessionToken string) (*domain.User, error)
}

type DataRepository interface {
	GetOneByUUID(uuid string) (*domain.Datas, error)
	Create(data *domain.Datas) (*domain.Datas, error)
	Delete(data *domain.Datas) error
	GetCountByType(dataType int, params domain.GetDataChartParams) (int64, error)
	GetOneByToken(token string) (*domain.User, error)
	GetDatasPaginate(in domain.GetDatasPaginateParams) ([]*domain.Datas, *domain.PaginateMeta, error)
}
