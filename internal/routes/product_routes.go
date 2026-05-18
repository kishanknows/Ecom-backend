package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kishanknows/Ecom-backend/internal/handlers"
)

func RegisterProductRoutes(
	r *gin.Engine,
){

	handler := handlers.NewProductHandler()

	r.GET(
		"/products",
		handler.GetProducts,
	)

	r.POST(
		"/products",
		handler.CreateProduct,
	)
}