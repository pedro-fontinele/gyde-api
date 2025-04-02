package controller

import (
	"gyde-api/model"
	"gyde-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController é responsável por gerenciar as requisições relacionadas a produtos.
type ProductController struct {
	productUseCase usecase.ProductUseCase
}

// Setup inicializa e retorna uma instância de ProductController.
func Setup(productUseCase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: productUseCase,
	}
}

// GetAll busca todos os produtos disponíveis e os retorna como resposta JSON.
func (productController *ProductController) GetAll(context *gin.Context) {
	products, err := productController.productUseCase.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	context.JSON(http.StatusOK, products)
}

// Create recebe os dados de um produto via JSON, valida e o insere no banco.
func (productController *ProductController) Create(context *gin.Context) {

	var product model.Product
	err := context.BindJSON(&product)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	createdProduct, err := productController.productUseCase.Create(product)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusCreated, createdProduct)
}

// GetById busca um produto pelo ID informado na URL.
func (productController *ProductController) GetById(context *gin.Context) {

	id := context.Param("id")
	if id == "" {
		response := model.Response{
			Message: "Id do produto nao pode ser nulo",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um numero",
		}
		context.JSON(http.StatusBadRequest, response)
		return
	}

	product, err := productController.productUseCase.GetById(productId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.Response{
			Message: "Produto nao foi encontrado na base de dados",
		}
		context.JSON(http.StatusNotFound, response)
		return
	}

	context.JSON(http.StatusOK, product)
}
