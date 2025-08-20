package main

/*
   	1:1 TECNICO COM O BRUNO LOPES
	-> Assunto: Closures
	Um conceito que vem do paradigma funcional e demonstra como o Go é uma linguagem multi-paradigma
	- Closure é uma função que captura variáveis do contexto em que foi definida
	- Podemos considerar a recursividade como um caso especial de closure, pois a função recursiva captura a si mesma

	EXEMPLO GETTER/SETTER
	- A função closure retorna duas funções: uma para setar um valor e outra para pegar um valor
	- A variável a é um map que é acessado pelas funções
	- A função getter retorna o valor da chave passada como argumento
	- A função setter seta o valor da chave passada como argumento
    - A variável só existe no contexto da função closure, mas as funções getter e setter podem acessá-la
	Isso é util para criar funções que encapsulam um estado interno e permitem manipulá-lo de forma controlada
*/

import "fmt"

type Setter func(name string, value string)
type Getter func(name string) string

func closure() (Getter, Setter) {
	a := map[string]string{}

	getter := func(name string) string {
		return a[name]
	}

	setter := func(name string, value string) {
		a[name] = value
	}

	return getter, setter
}

func main() {
	getter, setter := closure()

	setter("a", "1")

	fmt.Println(getter("a"))
}
