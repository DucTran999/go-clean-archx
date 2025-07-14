package usecase_test

import (
	"testing"

	"github.com/DucTran999/go-clean-archx/internal/dto"
	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/DucTran999/go-clean-archx/internal/port"
	"github.com/DucTran999/go-clean-archx/internal/usecase"
	"github.com/DucTran999/go-clean-archx/test/datatest"
	"github.com/DucTran999/go-clean-archx/test/mockbuilder"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		input       dto.CreateProductInput
		expectedErr error
		setupUT     func(t *testing.T) port.ProductUsecase // set up under test
	}{
		{
			name:  "success",
			input: dto.CreateProductInput{Name: "Book", Qty: 5, Price: 20},
			setupUT: func(t *testing.T) port.ProductUsecase {
				t.Helper()
				mRepo := mockbuilder.NewProductRepoBuilder(t).CreateProductSuccess().Build()
				return usecase.NewProductUsecase(mRepo)
			},
			expectedErr: nil,
		},
		{
			name:  "invalid product",
			input: dto.CreateProductInput{Name: "cool hat", Qty: 10, Price: 0},
			setupUT: func(t *testing.T) port.ProductUsecase {
				t.Helper()
				mRepo := mockbuilder.NewProductRepoBuilder(t).Build()
				return usecase.NewProductUsecase(mRepo)
			},
			expectedErr: entity.ErrInvalidPrice,
		},
		{
			name:  "failed cause db error",
			input: dto.CreateProductInput{Name: "Laptop", Qty: 10, Price: 999.99},
			setupUT: func(t *testing.T) port.ProductUsecase {
				t.Helper()
				mRepo := mockbuilder.NewProductRepoBuilder(t).CreateProductErrorDB().Build()
				return usecase.NewProductUsecase(mRepo)
			},
			expectedErr: datatest.ErrUnexpectedDB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Setup usecase under test
			uc := tt.setupUT(t)

			got, err := uc.CreateProduct(t.Context(), tt.input)

			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
				assert.Equal(t, tt.input.Name, got.Name)
				assert.Equal(t, tt.input.Qty, got.Qty)
				assert.Equal(t, tt.input.Price, got.Price)
				assert.Equal(t, datatest.FakeProductID, got.ID)
			}
		})
	}
}
