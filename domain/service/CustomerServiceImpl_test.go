package service

import (
	"testing"
	"ws_customers/domain/dto"
	"ws_customers/infrastructure/repository"

	"github.com/stretchr/testify/assert"
)

type fakeOrderRepository struct{}

var (
	createOrder  func(o dto.OrderDto) (interface{}, error)
	findOrder    func(id string) dto.OrderDto
	changeStatus func(idOrder string, status string) error
)

func (f *fakeOrderRepository) CreateOrder(o dto.OrderDto) (interface{}, error) {
	return createOrder(o)
}

func (f *fakeOrderRepository) FindOrder(id string) dto.OrderDto {
	return findOrder(id)
}

func (f *fakeOrderRepository) ChangeStatus(idOrder string, status string) error {
	return changeStatus(idOrder, status)
}

var orderRepositoryImplFake repository.OrderRepository = &fakeOrderRepository{}

func TestCreateOrders(t *testing.T) {

	createOrder = func(o dto.OrderDto) (interface{}, error) {
		return dto.OrderDto{
			CoordinatesDto: dto.CoordinatesDto{
				FromCoordinates: "24.53525235, 23.452355",
				ToCoordinates:   "24.53525235, 23.452355",
			},
		}, nil
	}

	order := &dto.OrderDto{
		CoordinatesDto: dto.CoordinatesDto{
			FromCoordinates: "24.53525235, 23.452355",
			ToCoordinates:   "24.53525235, 23.452355",
		},
	}

	id, err := orderRepositoryImplFake.CreateOrder(*order)

	assert.Nil(t, err)
	assert.NotNil(t, id)

}

func TestFindOrder(t *testing.T) {

	findOrder = func(id string) dto.OrderDto {
		return dto.OrderDto{
			Items: dto.Items{
				TotalItems: 3,
			},
		}
	}

	order := dto.OrderDto{
		Items: dto.Items{
			TotalItems: 3,
		},
	}

	orderResult := orderRepositoryImplFake.FindOrder("idTest")

	assert.NotNil(t, orderResult)
	assert.EqualValues(t, orderResult.TotalItems, order.TotalItems)

}
