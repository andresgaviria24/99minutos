package dto

import (
	"time"
	"ws_customers/constants"
)

type OrderDto struct {
	FromAddress AddressDto `json:"from_address"`
	ToAddress   AddressDto `json:"to_address"`
	CoordinatesDto
	Items
	Date   time.Time `json:"date_created,omitempty"`
	Status string    `json:"status,omitempty"`
}

func (o *OrderDto) InitOrder() OrderDto {
	o.Date = time.Now()
	o.Status = constants.Creado
	return *o
}
