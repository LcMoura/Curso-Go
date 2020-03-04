package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// Abrir a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:kja34913780@/cursogo")

	// Tratar o erro
	if err != nil {
		log.Fatal(err)
	}

	// Fechar a conexão com o banco de dados no final da função
	defer db.Close() // defer(ultima instrução)

	tx, _ := db.Begin() // Iniciarlizar uma transação
	stmt, _ := tx.Prepare("INSERT INTO usuarios(id, nome) VALUES(?, ?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	_, err = stmt.Exec(1, "Tiago") // chave duplicada
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
