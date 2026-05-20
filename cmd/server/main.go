package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kishanknows/Ecom-backend/internal/database"
	"github.com/kishanknows/Ecom-backend/internal/routes"
)

func main(){
	godotenv.Load("../../.env")

	r := gin.Default()
	database.Connect()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Powered by Go.",
		})
	})

	routes.RegisterProductRoutes(r)

	r.Run(":8080")
}