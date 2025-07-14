package entity_test

import (
	"testing"

	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/stretchr/testify/require"
)

func TestProduct_IsValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		product     entity.Product
		expectedErr error
	}{
		{
			name:        "valid product",
			product:     entity.Product{Name: "Laptop", Qty: 10, Price: 1200.5},
			expectedErr: nil,
		},
		{
			name:        "empty name",
			product:     entity.Product{Name: "", Qty: 10, Price: 1000},
			expectedErr: entity.ErrEmptyName,
		},
		{
			name:        "negative quantity",
			product:     entity.Product{Name: "Mouse", Qty: -5, Price: 25},
			expectedErr: entity.ErrQtyNegative,
		},
		{
			name:        "zero price",
			product:     entity.Product{Name: "Keyboard", Qty: 5, Price: 0},
			expectedErr: entity.ErrInvalidPrice,
		},
		{
			name:        "negative price",
			product:     entity.Product{Name: "Monitor", Qty: 3, Price: -200},
			expectedErr: entity.ErrInvalidPrice,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			err := tt.product.IsValid()

			require.Equal(t, err, tt.expectedErr)
		})
	}
}
