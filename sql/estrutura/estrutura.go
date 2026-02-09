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
	// abrindo a conexão com o banco de dados
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err)
	}
	// defer é responsável por adiar a execução de uma função até o final do escopo
	defer db.Close() // fecha a conexão com o banco ao final da função main

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
