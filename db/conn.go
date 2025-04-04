package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Configurações de conexão com o banco de dados PostgreSQL.
const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "gyde-api"
)

// ConnectDB estabelece uma conexão com o banco de dados PostgreSQL e retorna uma instância de sql.DB.
func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + dbname)

	return db, nil
}
