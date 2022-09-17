package service

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"ws_customers/constants"
	"ws_customers/domain/dto"
	address "ws_customers/domain/service/rules/addresses"
	"ws_customers/domain/service/rules/coordinates"
	"ws_customers/domain/service/rules/items"
	"ws_customers/infrastructure/persistence"
	"ws_customers/infrastructure/repository"
)

type CustomerServiceImpl struct {
	orderRepository repository.OrderRepository
}

func InitCustomertServiceImpl() *CustomerServiceImpl {
	dbHelper, err := persistence.InitDbHelper()
	if err != nil {
		log.Fatal(err.Error())
	}
	return &CustomerServiceImpl{
		orderRepository: dbHelper.OrderRepository,
	}
}

func (r *CustomerServiceImpl) CreateOrders(order dto.OrderDto) (interface{}, dto.Response) {

	var (
		w        sync.WaitGroup
		response dto.Response
	)

	c := make(chan error)
	w.Add(4)

	go func() {
		w.Wait()
		close(c)
	}()

	go func() {
		defer w.Done()

		err := coordinates.IsValidCoordinates(order.CoordinatesDto)
		if err != nil {
			c <- err
		}
	}()

	go func() {
		defer w.Done()

		err := address.IsValidAddress(order.FromAddress)
		if err != nil {
			c <- err
		}
	}()

	go func() {
		defer w.Done()

		err := address.IsValidAddress(order.ToAddress)
		if err != nil {
			c <- err
		}
	}()

	go func() {
		defer w.Done()

		err := items.IsValidItems(&order.Items)

		if err != nil {
			c <- err
		}
	}()

	for err := range c {
		if err != nil {
			response.Warnings = append(response.Warnings, err.Error())
		}
	}

	if len(response.Warnings) > 0 {
		response.Description = "Error"
		response.Status = http.StatusBadRequest
		return nil, response
	}

	result, err := r.orderRepository.CreateOrder(order.InitOrder())

	if err != nil {
		return nil, dto.Response{
			Status: http.StatusBadRequest,
		}
	}

	return result, dto.Response{
		Status: http.StatusOK,
	}
}

func (r *CustomerServiceImpl) FindOder(id string) interface{} {
	order := r.orderRepository.FindOrder(id)

	if len(order.FromCoordinates) > 0 {
		return order
	}

	return nil
}

func (r *CustomerServiceImpl) CancelOrder(id string) (error, dto.Response) {
	order := r.orderRepository.FindOrder(id)

	if len(order.FromCoordinates) == 0 {
		return errors.New("id not exist"), dto.Response{
			Status: http.StatusBadRequest,
		}
	}

	fmt.Println(order.Date)

	err := r.orderRepository.ChangeStatus(id, constants.Cancelado)

	if err != nil {
		return err, dto.Response{
			Status: http.StatusBadRequest,
		}
	}
	return nil, dto.Response{Status: http.StatusOK}
}
