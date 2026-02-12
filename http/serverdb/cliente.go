package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Usuário :)
type Usuario struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

// UsuarioHandler analisa o request e delega para a função correta
func UsuarioHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/") // tenta recuperar o id do user
	id, _ := strconv.Atoi(sid)                          // converte o id para inteiro, se for possível

	switch {
	case r.Method == "GET" && id > 0:
		usuarioPorId(w, r, id) // se for um GET e tiver um id válido, buscar o usuário
	case r.Method == "GET":
		usuarioTodos(w, r) // se for um GET sem id, buscar todos os usuários
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Rota não encontrada :(") // se for outro método, retornar 404
	}
}

func usuarioPorId(w http.ResponseWriter, r *http.Request, id int) {
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuario Usuario
	err = db.QueryRow("select id, nome from cursogo.usuarios where id = ?", id).Scan(&usuario.ID, &usuario.Nome)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Usuário não encontrado")
		return
	}

	json, _ := json.Marshal(usuario)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func usuarioTodos(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:pass123@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, _ := db.Query("select id, nome from cursogo.usuarios")
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome)
		usuarios = append(usuarios, u)
	}

	if len(usuarios) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Usuários não encontrados")
		return
	}

	json, _ := json.Marshal(usuarios)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
