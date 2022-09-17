package service

import (
	"log"
	"net/http"
	"ws_customers/domain/dto"
	"ws_customers/infrastructure/persistence"
	"ws_customers/infrastructure/repository"
)

type PrivateServiceImpl struct {
	orderRepository repository.OrderRepository
}

func InitPrivateServiceImpl() *PrivateServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &PrivateServiceImpl{
		orderRepository: dbHelper.OrderRepository,
	}
}

func (o *PrivateServiceImpl) ChangeStatus(id string, status string) dto.Response {

	err := o.orderRepository.ChangeStatus(id, status)

	if err != nil {
		return dto.Response{
			Status: http.StatusBadRequest,
		}
	}

	return dto.Response{
		Status: http.StatusOK,
	}
}
