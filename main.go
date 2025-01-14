package main

import (
	"api-go/controller"
	"api-go/db"
	"api-go/repository"
	"api-go/usecase"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//camada Usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/product", ProductController.CreateProduct)
	server.GET("/product/:productId", ProductController.GetProductById)

	server.Run(":8000")

}
