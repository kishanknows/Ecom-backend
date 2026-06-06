package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kishanknows/Ecom-backend/internal/database"
	"github.com/kishanknows/Ecom-backend/internal/metrics"
	"github.com/kishanknows/Ecom-backend/internal/middlewares"
	"github.com/kishanknows/Ecom-backend/internal/routes"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main(){
	godotenv.Load("../../.env")
	metrics.Register()

	r := gin.Default()
	r.Use(middlewares.MetricsMiddleware())
	
	database.Connect()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Powered by Go.",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	routes.RegisterProductRoutes(r)

	r.Run(":8080")
}