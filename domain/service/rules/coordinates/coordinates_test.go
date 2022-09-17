package coordinates

import (
	"errors"
	"testing"
	"ws_customers/domain/dto"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCoordinates(t *testing.T) {
	var cases = []struct {
		name        string
		coordinates dto.CoordinatesDto
		expected    error
	}{
		{
			name: "when is invalid coordinates is empty then return error",
			coordinates: dto.CoordinatesDto{
				FromCoordinates: "",
				ToCoordinates:   "",
			},
			expected: errors.New("coordinates wrong"),
		}, {
			name: "when is invalid coordinates with wrong format then return error",
			coordinates: dto.CoordinatesDto{
				FromCoordinates: "N23.43345, E32.6457",
				ToCoordinates:   "N23.43345, E32.6457",
			},
			expected: errors.New("coordinates wrong"),
		}, {
			name: "when is invalid coordinates without length then return error",
			coordinates: dto.CoordinatesDto{
				FromCoordinates: "23.43345",
				ToCoordinates:   "23.43345, E32.6457",
			},
			expected: errors.New("coordinates wrong"),
		}, {
			name: "when the latitude exceeds 90 then return error",
			coordinates: dto.CoordinatesDto{
				FromCoordinates: "20.43345, 195.43345",
				ToCoordinates:   "23.43345, 32.6457",
			},
			expected: errors.New("coordinates wrong"),
		}, {
			name: "when the length exceeds 180 then return error",
			coordinates: dto.CoordinatesDto{
				FromCoordinates: "80.43345, 190.43345",
				ToCoordinates:   "23.43345, 192.6457",
			},
			expected: errors.New("coordinates wrong"),
		}, {
			name: "when is valid coordinates then return nil",
			coordinates: dto.CoordinatesDto{
				FromCoordinates: "24.53525235, 23.452355",
				ToCoordinates:   "24.53525235, 23.45235",
			},
			expected: nil,
		}, {
			name: "when is valid coordinates into range then return nil",
			coordinates: dto.CoordinatesDto{
				FromCoordinates: "4, -3",
				ToCoordinates:   "4, -3",
			},
			expected: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			output := IsValidCoordinates(tt.coordinates)
			assert.EqualValues(t, output, tt.expected)
		})
	}
}
