package datasvc

import (
	"context"
	"errors"
	"log"

	"github.com/platformsh/template-golang/config/pusher"
	"github.com/platformsh/template-golang/domain"
	"github.com/platformsh/template-golang/repository"
	"github.com/platformsh/template-golang/service"
)

type dataService struct {
	dataRepo repository.DataRepository
}

func NewDataService(dataRepo repository.DataRepository) service.DataService {
	return &dataService{
		dataRepo: dataRepo,
	}
}

func (instance dataService) Create(ctx context.Context, in domain.CreateDataRequest) error {
	if _, err := instance.dataRepo.Create(in.ToData()); err != nil {
		return err
	}

	go func() {
		if datas, err := instance.GetTotalDataPerType(ctx, domain.GetDataChartParams{}); err != nil {
			log.Println("failed to get total data per type : ", err)
		} else {
			client := pusher.Authenticate()
			client.Trigger("mask_detector-development", "my-event", datas)
		}
	}()

	return nil
}

func (instance dataService) Delete(ctx context.Context, uuid string) error {
	// get data by uuid
	data, err := instance.dataRepo.GetOneByUUID(uuid)
	if err != nil {
		return err
	}

	if data.IsEmpty() {
		return errors.New("data not found")
	}

	if err := instance.dataRepo.Delete(data); err != nil {
		return errors.New("failed to delete data")
	}

	return nil
}

func (instance dataService) GetTotalDataPerType(ctx context.Context, in domain.GetDataChartParams) (*domain.DataChartTransfmer, error) {
	proper, err := instance.dataRepo.GetCountByType(1, in)
	if err != nil {
		return nil, errors.New("failed to count total of proper mask")
	}

	improper, err := instance.dataRepo.GetCountByType(2, in)
	if err != nil {
		return nil, errors.New("failed to count total of proper mask")
	}

	no, err := instance.dataRepo.GetCountByType(3, in)
	if err != nil {
		return nil, errors.New("failed to count total of proper mask")
	}

	return domain.NewDataChartTransfmer(int(proper), int(improper), int(no)), nil
}

func (instance dataService) GetDatasPaginate(ctx context.Context, in domain.GetDatasPaginateParams) (*domain.GetDataPaginateResponse, error) {
	datas, paginate, err := instance.dataRepo.GetDatasPaginate(in)
	if err != nil {
		return nil, err
	}

	var datasTransformer []domain.DataTransformer

	for _, data := range datas {
		datasTransformer = append(datasTransformer, *data.ToTransform())
	}

	return domain.NewGetDataPaginateResponse(datasTransformer, *paginate), nil
}
