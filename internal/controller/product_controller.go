package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/DucTran999/go-clean-archx/internal/dto"
	"github.com/DucTran999/go-clean-archx/internal/entity"
	"github.com/DucTran999/go-clean-archx/internal/port"
	"github.com/gin-gonic/gin"
)

// ProductController handles incoming HTTP requests and sends appropriate responses.
// It acts as the delivery layer in Clean Architecture, connecting HTTP routes to usecases.
type ProductController struct {
	productUC port.ProductUsecase
}

// NewProductController creates a new ProductController instance.
func NewProductController(productUC port.ProductUsecase) *ProductController {
	return &ProductController{
		productUC: productUC,
	}
}

// CreateProduct handles POST /products requests.
func (hdl *ProductController) CreateProduct(ctx *gin.Context) {
	var payload CreateProductRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		JSONBadRequestResponse(ctx, "invalid request payload", err)
		return
	}

	input := dto.CreateProductInput{
		Name:  payload.Name,
		Qty:   payload.Qty,
		Price: payload.Price,
	}

	created, err := hdl.productUC.CreateProduct(ctx.Request.Context(), input)
	if err != nil {
		if isClientError(err) {
			JSONBadRequestResponse(ctx, "validation failed", err)
		} else {
			log.Printf("[ERROR] op=create_product, err=%v", err)
			JSONInternalErrorResponse(ctx, "failed to create product")
		}
		return
	}

	JSONResponse(ctx, http.StatusCreated, APIResponse{
		Message: "product created successfully",
		Data:    gin.H{"id": created.ID.String()},
	})
}

func isClientError(err error) bool {
	return errors.Is(err, entity.ErrEmptyName) ||
		errors.Is(err, entity.ErrInvalidPrice) ||
		errors.Is(err, entity.ErrQtyNegative)
}
