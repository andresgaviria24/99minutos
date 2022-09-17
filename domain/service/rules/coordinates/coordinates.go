package coordinates

import (
	"errors"
	"strconv"
	"strings"
	"ws_customers/domain/dto"
)

func IsValidCoordinates(o dto.CoordinatesDto) error {

	if len(o.ToCoordinates) == 0 || len(o.FromCoordinates) == 0 {
		return errorMessage()
	}

	if !RuleCoordinates(o.FromCoordinates) || !RuleCoordinates(o.ToCoordinates) {
		return errorMessage()
	}

	return nil
}

func RuleCoordinates(o string) bool {
	coords := strings.Split(o, ", ")

	if len(coords) != 2 {
		return false
	}

	coord1, err1 := strconv.ParseFloat(coords[0], 64)
	coord2, err2 := strconv.ParseFloat(coords[1], 64)
	if err1 != nil || err2 != nil || coord1 < -90 || coord1 > 90 || coord2 < -180 || coord2 > 180 {
		return false
	}
	return true
}

func errorMessage() error {
	return errors.New("coordinates wrong")
}
