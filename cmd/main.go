package main

import (
	"gyde-api/controller"
	"gyde-api/db"
	"gyde-api/repository"
	"gyde-api/usecase"

	"github.com/gin-gonic/gin"
)

// main inicializa o servidor HTTP usando o framework Gin,
// configura a conexão com o banco de dados e define as rotas da API.
func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada repository
	ProductRepository := repository.Setup(dbConnection)

	// Camada usecase
	ProductUseCase := usecase.Setup(ProductRepository)

	// Camada controller
	ProductController := controller.Setup(ProductUseCase)

	server.GET("/product", ProductController.GetAll)
	server.POST("/product", ProductController.Create)
	server.GET("/product/:id", ProductController.GetById)

	server.Run(":8000")
}
