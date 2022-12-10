package service

import (
	"context"

	"github.com/platformsh/template-golang/domain"
)

type UserService interface {
	CreateNewUser(ctx context.Context, in domain.CreteUserRequest) error
	Login(ctx context.Context, in domain.LoginRequest) (sessionToken string, err error)
	Logout(ctx context.Context, in domain.LogoutRequest) error
	CheckToken(ctx context.Context, token string) (bool, error)
}

type DataService interface {
	Create(ctx context.Context, in domain.CreateDataRequest) error
	Delete(ctx context.Context, uuid string) error
	GetTotalDataPerType(ctx context.Context, in domain.GetDataChartParams) (*domain.DataChartTransfmer, error)
	GetDatasPaginate(ctx context.Context, in domain.GetDatasPaginateParams) (*domain.GetDataPaginateResponse, error)
}
