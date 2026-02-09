package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// preparando a query de update
	stmt, _ := db.Prepare("update cursogo.usuarios set nome = ? where id = ?")
	stmt.Exec("William Smith", 1)
	stmt.Exec("Mary Johnson", 2)

	// preparando a query de delete
	stmt2, _ := db.Prepare("delete from cursogo.usuarios where id = ?")
	stmt2.Exec(3)

}
