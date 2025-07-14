// Package controller contains HTTP delivery logic for the application.
// It maps incoming requests to use case calls and formats appropriate responses.
package controller

// CreateProductRequest defines the expected JSON structure for creating a new product.
// It includes validation rules using Gin's binding tags to enforce input correctness
// at the HTTP layer before data enters the application core.
type CreateProductRequest struct {
	Name  string  `json:"name" binding:"required"`
	Qty   int     `json:"qty"`
	Price float64 `json:"price"`
}
