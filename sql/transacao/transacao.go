package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// AULA -EXECUTANDO INSERÇÕES EM UMA TRANSAÇÃO

func main() {
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, _ := db.Begin() // iniciando a transação
	stmt, _ := tx.Prepare("insert into cursogo.usuarios (id, nome) values (?, ?)")

	stmt.Exec(2000, "Carlos")
	stmt.Exec(2001, "Fran ")
	_, err = stmt.Exec(2000, "Tiago") // erro de chave primária duplicada

	if err != nil {
		// no go não existe o conceito de exceção (try/catch)
		// então tratamos o erro explictamente

		tx.Rollback() // desfaz a transação em caso de erro
		log.Fatal(err)
	}

	tx.Commit() // confirma a transação
}
