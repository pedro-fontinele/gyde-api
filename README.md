# Gyde API

Gyde API é uma API desenvolvida em Go utilizando o framework Gin. Ela é containerizada com Docker e utiliza o banco de dados PostgreSQL. Sua principal responsabilidade é gerenciar o cadastro de produtos.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação utilizada para desenvolver a API.
- **Gin**: Framework web para Go, utilizado para facilitar o desenvolvimento.
- **Docker**: Utilizado para conteinerização da aplicação.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenamento dos produtos.

## Configuração e Execução

### Clonando o Repositório

```sh
git clone https://github.com/seu-usuario/gyde-api.git
cd gyde-api
```

### Configuração do Banco de Dados

Certifique-se de ter o PostgreSQL instalado e rodando. Altere as credenciais no arquivo de configuração da API caso necessário.

### Construindo e Executando com Docker

1. Construa a imagem do Docker:

```sh
docker-compose build
```

2. Suba os contêineres:

```sh
docker-compose up -d
```

A API estará rodando na porta especificada (por padrão, `:8080`).

## Endpoints Principais

### Criar Produto
```http
POST /product
```
**Body:**
```json
{
  "name": "Produto X",
  "price": 99.90
}
```

### Listar Produtos
```http
GET /product
```

### Buscar Produto por ID
```http
GET /product/{id}
```

## Contribuição

1. Fork este repositório.
2. Crie uma branch com sua feature: `git checkout -b minha-feature`.
3. Commit suas alterações: `git commit -m 'Minha nova feature'`.
4. Envie para o branch principal: `git push origin minha-feature`.
5. Abra um Pull Request.

## Licença

Este projeto está sob a licença MIT. Para mais detalhes, consulte o arquivo LICENSE.

