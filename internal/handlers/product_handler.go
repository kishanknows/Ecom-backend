package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kishanknows/Ecom-backend/internal/models"
	"github.com/kishanknows/Ecom-backend/internal/services"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{
		service: services.NewProductService(),
	}
}

func (h *ProductHandler) GetProducts(ctx *gin.Context){
	products, err := h.service.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(
		http.StatusOK,
		products,
	)
}

func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var product models.Product
	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	err = h.service.CreateProduct(&product)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		product,
	)
}