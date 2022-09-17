package address

import (
	"errors"
	"ws_customers/domain/dto"
)

func IsValidAddress(a dto.AddressDto) error {

	if len(a.Address) == 0 || len(a.Region) == 0 ||
		a.RegionId == 0 || len(a.ZipCode) == 0 {
		return errorMessage()
	}

	return nil
}

func errorMessage() error {
	return errors.New("address wrong")
}
