package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// preparando a query de insert
	stmt, _ := db.Prepare("insert into cursogo.usuarios (nome) values (?)")
	stmt.Exec("John Doe")
	stmt.Exec("Jane Smith")

	// a partir da resposta temos metodos que podemos verificar
	res, _ := stmt.Exec("Alice Johnson")

	// ultimo id inserido
	id, _ := res.LastInsertId()
	fmt.Println(id)

	// quantas linhas foram afetadas
	linhas, _ := res.RowsAffected()
	fmt.Println(linhas)

}
