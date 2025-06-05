package main

import "fmt"

/*
	AULA - EXPLORANDO COMANDOS
	- go run [arquivo]: compila e executa o arquivo Go especificado
	- go help [comando]: exibe ajuda específica para um comando
	- go version: exibe a versão do Go instalada
	- godoc -http: inicia um servidor local para acessar a documentação do Go
	- go env: exibe as variáveis de ambiente do Go
	- go vet [arquivo]: analisa o código em busca de erros comuns
	- go build [arquivo]: compila o código Go, gerando um executável

	PARA INSTALAR PACOTES:
	- go mod init [nome_do_modulo]: inicializa um novo módulo Go
	- go mod tidy: remove dependências não utilizadas e atualiza o go.mod
	- go get -u [pacote]: baixa e atualiza um pacote para a versão mais recente

*/

func main() {
	fmt.Printf("Outro programa em %s\n", "Go")
}
