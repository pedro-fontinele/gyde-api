package usecase

import (
	"gyde-api/model"
	"gyde-api/repository"
)

type ProductUseCase struct {
	productRepository repository.ProductRepository
}

func Setup(productRepository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		productRepository: productRepository,
	}
}

func (productUseCase *ProductUseCase) GetAll() ([]model.Product, error) {
	return productUseCase.productRepository.GetAll()
}

func (productUseCase *ProductUseCase) Create(product model.Product) (model.Product, error) {

	productId, err := productUseCase.productRepository.Create(product)
	if err != nil {
		return model.Product{}, err
	}

	product.Id = productId

	return product, nil
}

func (productUseCase *ProductUseCase) GetById(id int) (*model.Product, error) {

	product, err := productUseCase.productRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
