package main

// GERADO PELO GEMINI

import (
	"fmt"
	"math/rand"
	"time"
)

// Constantes para o tamanho da cartela
const (
	TamanhoCartela = 5
	MaxNumeroBingo = 75 // Números de 1 a 75
)

// Função para gerar uma cartela de Bingo
// Retorna um slice de slices (matriz) representando a cartela
func gerarCartelaBingo() [][]int {
	cartela := make([][]int, TamanhoCartela) // Cria as 5 linhas
	numerosUsados := make(map[int]bool)      // Ajuda a garantir números únicos

	for i := 0; i < TamanhoCartela; i++ {
		cartela[i] = make([]int, TamanhoCartela) // Cria as 5 colunas para cada linha
		for j := 0; j < TamanhoCartela; j++ {
			// Gera números aleatórios únicos para a cartela
			numero := 0
			for {
				numero = rand.Intn(MaxNumeroBingo) + 1 // Gera de 1 a 75
				if !numerosUsados[numero] {
					numerosUsados[numero] = true
					break
				}
			}
			cartela[i][j] = numero
		}
	}
	return cartela
}

// Função para imprimir a cartela de forma formatada
func imprimirCartela(cartela [][]int, numerosSorteados map[int]bool) {
	fmt.Println("\n--- SUA CARTELA DE BINGO ---")
	for i := 0; i < TamanhoCartela; i++ {
		for j := 0; j < TamanhoCartela; j++ {
			num := cartela[i][j]
			if numerosSorteados[num] {
				// Se o número foi sorteado, imprime 'X' (ou um espaço)
				// Para visualização, vamos usar **X**
				fmt.Printf("%4s ", "X")
			} else {
				fmt.Printf("%4d ", num)
			}
		}
		fmt.Println()
	}
	fmt.Println("---------------------------\n")
}

// Função para verificar se o jogador ganhou (linha ou coluna)
func verificarVitoria(cartela [][]int, numerosSorteados map[int]bool) bool {
	// Verificar linhas
	for i := 0; i < TamanhoCartela; i++ {
		linhaCompleta := true
		for j := 0; j < TamanhoCartela; j++ {
			if !numerosSorteados[cartela[i][j]] {
				linhaCompleta = false
				break
			}
		}
		if linhaCompleta {
			fmt.Println("BINGO! LINHA COMPLETA!")
			return true
		}
	}

	// Verificar colunas
	for j := 0; j < TamanhoCartela; j++ {
		colunaCompleta := true
		for i := 0; i < TamanhoCartela; i++ {
			if !numerosSorteados[cartela[i][j]] {
				colunaCompleta = false
				break
			}
		}
		if colunaCompleta {
			fmt.Println("BINGO! COLUNA COMPLETA!")
			return true
		}
	}

	// Poderíamos adicionar verificação de diagonais aqui também
	// Por simplicidade, vamos focar em linhas e colunas para este exercício.

	// diagonal = j = i
	diagonalPrincipalCompleta := true
	for i := 0; i < TamanhoCartela; i++ {
		if !numerosSorteados[cartela[i][i]] {
			diagonalPrincipalCompleta = false
			break
		}
	}
	if diagonalPrincipalCompleta {
		fmt.Println("BINGO! DIAGONAL PRINCIPAL COMPLETA!")
		return true
	}

	// diagonal secundária
	diagonalSecundariaCompleta := true
	for i := 0; i < TamanhoCartela; i++ {
		j := (TamanhoCartela - 1) - i
		if !numerosSorteados[cartela[i][j]] {
			diagonalSecundariaCompleta = false
			break
		}
	}
	if diagonalSecundariaCompleta {
		fmt.Println("BINGO! DIAGONAL SECUNDÁRIA COMPLETA!")
		return true
	}

	return false
}

func main() {
	// Inicializa o gerador de números aleatórios
	rand.Seed(time.Now().UnixNano())

	// 1. Gerar a cartela do jogador
	cartelaJogador := gerarCartelaBingo()

	// Slice para armazenar os números que já foram sorteados
	// Usamos um map para busca rápida e para controlar números únicos sorteados
	numerosSorteados := make(map[int]bool)

	// Slice para os números que ainda podem ser sorteados
	poolNumeros := make([]int, 0, MaxNumeroBingo)
	for i := 1; i <= MaxNumeroBingo; i++ {
		poolNumeros = append(poolNumeros, i)
	}
	// Embaralha o pool de números para simular o sorteio aleatório
	rand.Shuffle(len(poolNumeros), func(i, j int) {
		poolNumeros[i], poolNumeros[j] = poolNumeros[j], poolNumeros[i]
	})

	fmt.Println("Bem-vindo ao Bingo Simplificado!")
	fmt.Println("Pressione Enter para sortear o próximo número...")

	jogadorGanhou := false
	numeroSorteioAtual := 0

	//for i := 0; i < TamanhoCartela; i++ {
	//	numerosSorteados[cartelaJogador[i][i]] = true
	//}
	//fmt.Println("MODO DE TESTE: FORÇANDO VITÓRIA NA DIAGONAL PRINCIPAL!")

	for i := 0; i < TamanhoCartela; i++ {
		j := (TamanhoCartela - 1) - i
		numerosSorteados[cartelaJogador[i][j]] = true
	}
	fmt.Println("MODO DE TESTE: FORÇANDO VITÓRIA NA DIAGONAL SECUNDARIA!")

	// Loop principal do jogo
	for !jogadorGanhou && numeroSorteioAtual < len(poolNumeros) {
		// Espera pelo Enter para continuar
		fmt.Scanln()

		// Sortear um número
		sorteado := poolNumeros[numeroSorteioAtual]

		if _, exists := numerosSorteados[poolNumeros[numeroSorteioAtual]]; !exists {
			sorteado := poolNumeros[numeroSorteioAtual]
			numerosSorteados[sorteado] = true // Adiciona ao map de sorteados
			fmt.Printf("Número sorteado: %d\n", sorteado)
		} else {
			// Se o número já está nos sorteados (por exemplo, da diagonal forçada),
			// apenas avança para o próximo.
			fmt.Printf("Número %d já foi sorteado (ou forçado para teste).\n", poolNumeros[numeroSorteioAtual])
		}
		numeroSorteioAtual++
		numerosSorteados[sorteado] = true // Adiciona ao map de sorteados

		//fmt.Printf("Número sorteado: %d\n", sorteado)
		imprimirCartela(cartelaJogador, numerosSorteados)

		// Verificar se o jogador ganhou
		jogadorGanhou = verificarVitoria(cartelaJogador, numerosSorteados)

		if !jogadorGanhou && numeroSorteioAtual == len(poolNumeros) {
			fmt.Println("Todos os números foram sorteados. Fim de jogo! Ninguém ganhou.")
		}
	}

	if jogadorGanhou {
		fmt.Println("\nParabéns! Você fez BINGO!")
	} else {
		fmt.Println("\nGame Over. Tente novamente!")
	}
}
