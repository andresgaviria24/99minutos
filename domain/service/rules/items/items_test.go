package items

import (
	"errors"
	"testing"
	"ws_customers/domain/dto"

	"github.com/stretchr/testify/assert"
)

func TestIsValidItems(t *testing.T) {
	var cases = []struct {
		name     string
		items    dto.Items
		size     string
		expected error
	}{
		{
			name: "when the numbers items are diferent then return error",
			items: dto.Items{
				TotalItems: 3,
				Items: []dto.Item{
					{
						Weight: 1,
						Type:   "",
					},
				},
			},
			expected: errors.New("items number wrong"),
		}, {
			name: "when the numbers items is ok then return nil",
			items: dto.Items{
				TotalItems: 1,
				Items: []dto.Item{
					{
						Weight: 10,
						Type:   "",
					},
				},
			},
			size:     "M",
			expected: nil,
		}, {
			name: "when the weight is 10 then return size M",
			items: dto.Items{
				TotalItems: 1,
				Items: []dto.Item{
					{
						Weight: 10,
						Type:   "",
					},
				},
			},
			size:     "M",
			expected: nil,
		}, {
			name: "when the weight is 5 then return size S",
			items: dto.Items{
				TotalItems: 1,
				Items: []dto.Item{
					{
						Weight: 5,
						Type:   "",
					},
				},
			},
			size:     "S",
			expected: nil,
		}, {
			name: "when the weight is 20 then return size L",
			items: dto.Items{
				TotalItems: 1,
				Items: []dto.Item{
					{
						Weight: 20,
						Type:   "",
					},
				},
			},
			size:     "L",
			expected: nil,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {

			output := IsValidItems(&tt.items)
			assert.EqualValues(t, output, tt.expected)
			assert.EqualValues(t, tt.items.Items[0].Type, tt.size)
		})
	}
}
