package service

import "ws_customers/domain/dto"

type OrderService interface {
	CreateOrders(dto.OrderDto) (interface{}, dto.Response)
	FindOder(id string) interface{}
	CancelOrder(id string) (error, dto.Response)
}
