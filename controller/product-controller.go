package controller

import (
	"gyde-api/model"
	"gyde-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func Setup(productUseCase usecase.ProductUseCase) ProductController {
	return ProductController{
		productUseCase: productUseCase,
	}
}

func (productController *ProductController) GetAll(context *gin.Context) {
	products, err := productController.productUseCase.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
	}

	context.JSON(http.StatusOK, products)
}

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
