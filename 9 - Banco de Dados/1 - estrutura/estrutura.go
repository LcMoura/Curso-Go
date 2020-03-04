package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// Conexão com o banco de dados
	db, err := sql.Open("mysql", "root:kja34913780@/")

	// Tratar o erro
	if err != nil {
		panic(err) // parar o programa
	}
	defer db.Close() // defer(ultima instrução) - fecha a conexão com o banco de dados

	// Executa as instruções SQL
	exec(db, "CREATE DATABASE IF NOT EXISTS cursogo")
	exec(db, "USE cursogo")
	exec(db, "DROP TABLE IF EXISTS usuarios")
	exec(db, `CREATE TABLE usuarios (
		id INTEGER AUTO_INCREMENT,
		nome VARCHAR(80),
		PRIMARY KEY(id)
	)`)
}
