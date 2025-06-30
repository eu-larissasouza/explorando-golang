package main

import "fmt"

// ESTRUTURAS DE DADOS EM GOLANG
// ARTIGO DE LEITURA: https://lucasmagnum.medium.com/iniciando-em-go-estrutura-de-dados-c0e71cc9146a

func main() {
	// --------------------------------------------------------------------------------
	// ESTRUTURA DE DADOS - ARRAY
	// Possuem um tamanho máximo de elementos que deve ser informado no momento em que o array for declarado

	// Declaramos um array do tipo inteiro com 10 posições
	var numeros [10]int
	// Todas as posições do array foram preenchidas com o valor zero
	fmt.Println(numeros) // --> [0 0 0 0 0 0 0 0 0 0]

	// COMO UM ARRAY PODE SER DECLARADO:

	// Declaramos um array de tamanho 5, todas as posições foram inicializadas para o 'zero value' do tipo int
	var exemplo1 [5]int
	fmt.Println("Exemplo 1", exemplo1) // --> [0 0 0 0 0]

	// Declaração curta com o array, o comportamento é o mesmo da linha anterior, precisamos inicializar o array
	// já que estamos usando a versão reduzida de declaração. Para inicializar um array utilizamos chaves.
	exemplo2 := [5]int{}
	fmt.Println("Exemplo 2", exemplo2) // --> [0 0 0 0 0]

	// Declaração com inicialização do array
	exemplo3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Exemplo 3", exemplo3) // --> [1 2 3 4 5]

	// Declaração com inicialização incompleta do array
	exemplo4 := [5]int{1, 2, 3}
	fmt.Println("Exemplo 4", exemplo4) // --> [1 2 3 0 0]

	// Declaração com inicialização de elementos específicos
	// Nesse caso inicializamos os elementos de index 0, 2 e 4
	exemplo5 := [5]int{0: 1, 2: 3, 4: 5}
	fmt.Println("Exemplo 5", exemplo5) // --> [1 0 3 0 5]

	// COMO ACESSAR ELEMENTOS DO ARRAY
	// zero-based numbering

	numerosStr := [5]string{"um", "dois", "três", "quatro", "cinco"}
	fmt.Println(numerosStr[0]) //-> "um"
	fmt.Println(numerosStr[1]) // -> "dois"
	fmt.Println(numerosStr[2]) // -> "três"
	fmt.Println(numerosStr[3]) // -> "quatro"
	fmt.Println(numerosStr[4]) // -> "cinco"

	fmt.Println("Alterando valores de index:")
	numerosStr = [5]string{"zero", "um", "dois", "quatro", "três"}
	fmt.Println(numerosStr[3]) // --> "quatro"
	fmt.Println(numerosStr[4]) // --> "três"

	numerosStr[3] = "três"
	numerosStr[4] = "quatro"

	fmt.Println(numerosStr[3]) // --> "três"
	fmt.Println(numerosStr[4]) // --> "quatro"
	// --------------------------------------------------------------------------------

	// ESTRUTURA DE DADOS SLICE
	// Conceito de arrays dinâmicos que podem aumentar ou diminuir de tamanho, se necessário
	// -> Utiliza um array internamente
	// -> Possui 3 atributos: ponteiro para o array, tamanho atual do slice e capacidade do array.

	meuSlice := make([]int, 3, 5)
	fmt.Println(meuSlice)

	// COMO DECLARAR UM SLICE
	meuSlice = make([]int, 0)    // Slice tamanho 0
	meuSlice = make([]int, 0, 4) // Slice tamanho 0 e capacidade 4
	meuSlice = make([]int, 5)    // Slice tamanho 5

	// DECLARAÇÃO CURTA: -------

	// SLICE NULO: Valor zero de um slice: nil
	// O slice precisa apontar para algum array e como não foi atribuído valores, logo, ele recebe esse valor nulo.

	var meuSlice2 []int // Slice nulo
	fmt.Println(meuSlice2)

	// SLICE VAZIO: Não possui nenhum valor dentro dele, mas o seu array já foi criado.
	meuSlice = []int{}        // Slice vazio
	meuSlice = make([]int, 0) // Slice vazio

	// COMO EXPANDIR O TAMANHO DO SLICE
	// Slices são dinâmicos, ao adicionar um novo item, ele armazena-o e internamente ele aumenta a sua capacidade.

	// Slice de tamanho 0 e capacidade 2
	meuSlice = make([]int, 0, 2) // ---> []
	// Adicionando um elemento ao slice
	meuSlice = append(meuSlice, 1) // ---> [1]
	// Adicionando outro elemento ao slice
	meuSlice = append(meuSlice, 2) // ---> [1 2]

	fmt.Println("Tamanho:", len(meuSlice))    //- -> Tamanho: 2
	fmt.Println("Capacidade:", cap(meuSlice)) // --> Capacidade: 2

	// Adicionando o terceiro elemento ao slice
	meuSlice = append(meuSlice, 3) // --> [1 2 3]

	fmt.Println("Tamanho:", len(meuSlice))    // --> Tamanho: 3
	fmt.Println("Capacidade:", cap(meuSlice)) // --> Capacidade: 4

	// COMO UM SLICE AUMENTA SUA CAPACIDADE?

	// Sempre que o array utilizado pelo slice estoura o seu limite de armazenamento, é internamente
	// feito um processo de expansão. Um novo array é criado com o dobro da capacidade anterior e todos
	// os elementos são copiados para o novo array e o slice apontado para o array com a capacidade maior.

	// OBSERVAÇÃO: Esse é um detalhe importante, pois podemos definir o tamanho do array que será criado
	// inicialmente e poupar com que o processo de criação e cópia do array aconteça.

	// --------------------------------------------------------------------------------

}
