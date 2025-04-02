package usecase

import (
	"gyde-api/model"
	"gyde-api/repository"
)

// ProductUseCase contém a lógica de negócio para manipulação de produtos.
type ProductUseCase struct {
	productRepository repository.ProductRepository
}

// Setup inicializa e retorna uma instância de ProductUseCase com a camada de repositório.
func Setup(productRepository repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		productRepository: productRepository,
	}
}

// GetAll busca todos os produtos através do repositório e os retorna.
func (productUseCase *ProductUseCase) GetAll() ([]model.Product, error) {
	return productUseCase.productRepository.GetAll()
}

// Create insere um novo produto, obtém seu ID e o retorna com os dados atualizados.
func (productUseCase *ProductUseCase) Create(product model.Product) (model.Product, error) {

	productId, err := productUseCase.productRepository.Create(product)
	if err != nil {
		return model.Product{}, err
	}

	product.Id = productId

	return product, nil
}

// GetById busca um produto pelo seu ID através do repositório e o retorna.
func (productUseCase *ProductUseCase) GetById(id int) (*model.Product, error) {

	product, err := productUseCase.productRepository.GetById(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}
