package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDataBase() *sql.DB {
	// Estrutura da string de conexão com o DB = "user dbname password host sslmode"
	connectString := "user=postgres dbname=loja password=adv17667 host=localhost port=5432 sslmode=disable"

	// Abrindo conexão com a base de dados
	db, err := sql.Open("postgres", connectString)
	if err != nil {
		panic(err)
	}

	return db
}
