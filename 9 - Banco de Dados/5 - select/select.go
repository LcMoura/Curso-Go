package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// criar struct para mapeamento
type usuario struct {
	id   int
	nome string
}

func main() {
	// Abrir a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:kja34913780@/cursogo")

	// Tratar o erro
	if err != nil {
		log.Fatal(err) // Registro de log do erro
	}
	defer db.Close() // Fechar a conexão com o banco de dados ao final da função

	// Realizar consulta no banco - SELECT
	rows, _ := db.Query("SELECT id, nome FROM usuarios WHERE id > ?", 5)
	defer rows.Close() // Fechar a consulta da linha no banco de dados

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}
}
