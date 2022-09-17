package repository

import "ws_customers/domain/dto"

type OrderRepository interface {
	CreateOrder(o dto.OrderDto) (interface{}, error)
	FindOrder(id string) dto.OrderDto
	ChangeStatus(idOrder string, status string) error
}
