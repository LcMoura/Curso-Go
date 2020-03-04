package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Abrir a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:kja34913780@/cursogo")

	// Tratar o erro
	if err != nil {
		panic(err) // parar o programa
	}

	// Fechar a conexão com o banco de dados no final da função
	defer db.Close() // defer(ultima instrução)

	// Statement - prepara a conexão com o banco
	stmt, _ := db.Prepare("INSERT INTO usuarios(nome) VALUES(?)")

	// Executa as instruções SQL
	stmt.Exec("Maria")
	stmt.Exec("João")
	stmt.Exec("Luiz")

	// Pegar o valor do ultimo ID
	res, _ := stmt.Exec("Vinícius")
	id, _ := res.LastInsertId()
	fmt.Println(id)

	// Pagar a quantidade de linhas afetadas
	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)
}
