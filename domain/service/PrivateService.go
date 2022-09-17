package service

import "ws_customers/domain/dto"

type PrivateService interface {
	ChangeStatus(string, string) dto.Response
}
