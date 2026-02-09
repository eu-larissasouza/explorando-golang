package main

// AULA - EXECUTANDO SELECT E MAPEANDO PARA STRUCTS

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Usuario struct {
	ID   int
	Nome string
}

func main() {
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// realizando o select
	rows, err := db.Query("select id, nome from cursogo.usuarios where id > ?", 5)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// enquanto houver linhas
	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome)
		fmt.Println(u)
	}
}
