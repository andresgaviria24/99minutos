package address

import (
	"errors"
	"testing"
	"ws_customers/domain/dto"

	"github.com/stretchr/testify/assert"
)

func TestIsValidAddress(t *testing.T) {
	var cases = []struct {
		name     string
		address  dto.AddressDto
		expected error
	}{
		{
			name: "when is invalid total address then return error",
			address: dto.AddressDto{
				Address:  "",
				ZipCode:  "",
				Region:   "",
				RegionId: 0,
			},
			expected: errors.New("address wrong"),
		}, {
			name: "when is invalid address then return error",
			address: dto.AddressDto{
				Address:  "",
				ZipCode:  "10394",
				Region:   "Test",
				RegionId: 1,
			},
			expected: errors.New("address wrong"),
		},
		{
			name: "when is invalid address then return error",
			address: dto.AddressDto{
				Address:  "Test 1",
				ZipCode:  "7478",
				Region:   "Region Test",
				RegionId: 1,
			},
			expected: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			output := IsValidAddress(tt.address)
			assert.EqualValues(t, output, tt.expected)
		})
	}
}
