package entity_test

import (
	"testing"

	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/stretchr/testify/require"
)

// In Go, test functions must start with the keyword `Test` and be placed in
// files ending with `_test.go`.
//
// The test file should be located next to the business logic file for easy
// lookup. For example, if your source file is `product.go`, the test file
// should be `product_test.go`.
//
// This layout helps Go tools detect tests for execution and code coverage.
func TestProduct_IsValid(t *testing.T) {
	t.Parallel() // This test can run in parallel with other test functions

	// Define the test case structure
	type testcase struct {
		name        string
		product     entity.Product
		expectedErr error
	}

	// Prepare test cases
	testTable := []testcase{
		{
			name:        "valid product",
			product:     entity.Product{Name: "Laptop", Qty: 10, Price: 1200.5},
			expectedErr: nil,
		},
		{
			name:        "empty name",
			product:     entity.Product{Name: "", Qty: 10, Price: 1000},
			expectedErr: entity.ErrProductInvalid,
		},
		{
			name:        "negative quantity",
			product:     entity.Product{Name: "Mouse", Qty: -5, Price: 25},
			expectedErr: entity.ErrProductInvalid,
		},
		{
			name:        "zero price",
			product:     entity.Product{Name: "Keyboard", Qty: 5, Price: 0},
			expectedErr: entity.ErrProductInvalid,
		},
		{
			name:        "negative price",
			product:     entity.Product{Name: "Monitor", Qty: 3, Price: -200},
			expectedErr: entity.ErrProductInvalid,
		},
	}

	// Run each test case
	for _, tc := range testTable {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel() // Run subtests in parallel

			// Act: call the method being tested
			err := tc.product.IsValid()

			// Assert: check the result against the expected error
			require.Equal(t, err, tc.expectedErr)
		})
	}
}
