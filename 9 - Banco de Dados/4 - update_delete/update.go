package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	/*
		Iniciar conexão
		Fechar conexão

		Realizar as transações
	*/

	// Abrir a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:kja34913780@/cursogo")

	// Tratar o erro
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// UPDATE
	stmt, _ := db.Prepare("UPDATE usuarios SET nome = ? WHERE id = ?")
	stmt.Exec("Luiz Carlos", 1)
	stmt.Exec("Luiz Vinícius", 2)

	//DELETE
	stmt2, _ := db.Prepare("DELETE * FROM usuarios WHERE id = ?")
	stmt2.Exec(3)
}
