package items

import (
	"errors"
	"ws_customers/constants"
	"ws_customers/domain/dto"
)

func IsValidItems(i *dto.Items) error {

	if i.TotalItems != len(i.Items) {
		return errors.New("items number wrong")
	}

	for n, v := range i.Items {
		switch {
		case v.Weight <= 5:
			i.Items[n].Type = constants.S
		case v.Weight <= 15:
			i.Items[n].Type = constants.M
		case v.Weight <= 25:
			i.Items[n].Type = constants.L
		default:
			return errors.New("does not have the standard service and must contact " +
				"the company to make an agreement special")
		}
	}

	return nil
}
