package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DucTran999/go-clean-archx/internal/controller"
	"github.com/DucTran999/go-clean-archx/test/mockbuilder"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProductController_CreateProduct(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		setupPayload   func(t *testing.T) []byte
		setupUT        func(t *testing.T) *controller.ProductController
		expectedStatus int
	}{
		{
			name: "success",
			setupPayload: func(t *testing.T) []byte {
				t.Helper()
				body := map[string]any{
					"name":  "Test Product",
					"qty":   10,
					"price": 99.99,
				}
				payload, err := json.Marshal(body)
				require.NoError(t, err)
				return payload
			},
			setupUT: func(t *testing.T) *controller.ProductController {
				t.Helper()
				productUC := mockbuilder.NewProductUsecaseBuilder(t).CreateProductSuccess().Build()
				return controller.NewProductController(productUC)
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name: "validation error from usecase",
			setupPayload: func(t *testing.T) []byte {
				t.Helper()
				body := map[string]any{
					"name":  "Test Product",
					"qty":   10,
					"price": 0,
				}
				payload, err := json.Marshal(body)
				require.NoError(t, err)
				return payload
			},
			setupUT: func(t *testing.T) *controller.ProductController {
				t.Helper()
				productUC := mockbuilder.NewProductUsecaseBuilder(t).CreateProductReturnsInvalidPrice().Build()
				return controller.NewProductController(productUC)
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid request payload",
			setupPayload: func(t *testing.T) []byte {
				t.Helper()
				body := map[string]any{
					"qty":   10,
					"price": 0,
				}
				payload, err := json.Marshal(body)
				require.NoError(t, err)
				return payload
			},
			setupUT: func(t *testing.T) *controller.ProductController {
				t.Helper()
				productUC := mockbuilder.NewProductUsecaseBuilder(t).Build()
				return controller.NewProductController(productUC)
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "unexpected error",
			setupPayload: func(t *testing.T) []byte {
				t.Helper()
				body := map[string]any{
					"name":  "Test Product",
					"qty":   10,
					"price": 10.2,
				}
				payload, err := json.Marshal(body)
				require.NoError(t, err)
				return payload
			},
			setupUT: func(t *testing.T) *controller.ProductController {
				t.Helper()
				productUC := mockbuilder.NewProductUsecaseBuilder(t).CreateProductReturnErrDB().Build()
				return controller.NewProductController(productUC)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Arrange
			controller := tt.setupUT(t)

			// Setup Gin router
			r := gin.Default()
			r.POST("/products", controller.CreateProduct)
			reqBody := tt.setupPayload(t)

			req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			// Act
			r.ServeHTTP(resp, req)

			// Assert
			assert.Equal(t, tt.expectedStatus, resp.Code)
		})
	}
}
