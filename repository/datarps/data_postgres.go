package datarps

import (
	"errors"

	"github.com/platformsh/template-golang/domain"
	"github.com/platformsh/template-golang/repository"
	"gorm.io/gorm"
)

type dataRepo struct {
	maria *gorm.DB
}

func NewDataRepository(maria *gorm.DB) repository.DataRepository {
	return &dataRepo{maria: maria}
}

func (instance dataRepo) GetOneByUUID(uuid string) (*domain.Datas, error) {
	var data *domain.Datas

	if err := instance.maria.Where("uuid = ?", uuid).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

func (instance dataRepo) Create(data *domain.Datas) (*domain.Datas, error) {
	if err := instance.maria.Save(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (instance dataRepo) Delete(data *domain.Datas) error {
	if err := instance.maria.Where("uuid = ?", data.UUID).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

func (instance dataRepo) GetCountByType(dataType int, params domain.GetDataChartParams) (int64, error) {
	var (
		count int64
		datas *domain.Datas
	)

	q := instance.maria.Where("mask_type = ?", dataType).Debug()

	if params.Start != "" {
		q = q.Where("created_at > ?", params.Start)
	}
	if params.End != "" {
		q = q.Where("created_at < ?", params.End)
	}

	if err := q.Model(&datas).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (instance dataRepo) GetOneByToken(token string) (*domain.User, error) {
	var user *domain.User

	if err := instance.maria.Where("session_token = ?", token).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (instance dataRepo) GetDatasPaginate(in domain.GetDatasPaginateParams) ([]*domain.Datas, *domain.PaginateMeta, error) {
	var (
		datas      []*domain.Datas
		totalDatas int64
		model      *domain.Datas
	)
	q := instance.maria.Debug()

	if !in.IsTypeEmpty() {
		q = q.Where("mask_type = ?", *in.Type)
	}
	if !in.IsStartEmpty() {
		q = q.Where("created_at >= ?", *in.Start)
	}
	if !in.IsEndEmpty() {
		q = q.Where("created_at <= ?", *in.End)
	}

	// getting total datas
	if err := q.Model(&model).Count(&totalDatas).Error; err != nil {
		return nil, nil, err
	}

	// get datas
	offset := (in.Page - 1)
	if err := q.Debug().Offset(offset).Limit(in.Limit).Order("created_at desc").Find(&datas).Error; err != nil {
		return nil, nil, err
	}

	totalPage := (int(totalDatas) / in.Limit) + (int(totalDatas) % in.Limit)

	return datas, domain.NewPaginateMeta(in.Page, totalPage, int(totalDatas)), nil
}
