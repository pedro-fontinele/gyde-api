package repository

import (
	"database/sql"
	"fmt"
	"gyde-api/model"
)

// ProductRepository é responsável por gerenciar a comunicação com a tabela de produtos no banco de dados.
type ProductRepository struct {
	connection *sql.DB
}

// Setup inicializa e retorna uma instância de ProductRepository com a conexão ao banco de dados.
func Setup(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

// GetAll busca todos os produtos armazenados no banco de dados.
func (pr *ProductRepository) GetAll() ([]model.Product, error) {
	rows, err := pr.connection.Query("select id, name, price from product")
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.Id,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}

// Create insere um novo produto no banco de dados e retorna o ID gerado.
func (pr *ProductRepository) Create(product model.Product) (int, error) {
	query, err := pr.connection.Prepare("INSERT INTO product (name, price) VALUES ($1, $2) RETURNING Id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&product.Id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return product.Id, nil
}

// GetById busca um produto pelo seu ID no banco de dados.
func (pr *ProductRepository) GetById(id int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id).Scan(
		&produto.Id,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &produto, nil
}
