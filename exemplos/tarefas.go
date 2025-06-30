package main

import (
	"fmt"
	"strconv"
	"strings"
)

// taskList é uma variável global que armazena todas as tarefas.
var taskList = make([]Task, 0)

// Status define o tipo para o estado da tarefa.
type Status uint8

// Constantes para cada um dos status, usando iota para atribuição sequencial.
const (
	ToDo       = iota // 0: Tarefa a fazer
	InProgress        // 1: Tarefa em progresso
	Done              // 2: Tarefa concluída
)

// Stringer para o tipo Status.
// Este metodo é automaticamente chamado por fmt.Print, fmt.Println e fmt.Printf (%s, %v)
// para obter a representação em string do status.
func (s Status) String() string {
	switch s {
	case ToDo:
		return "A Fazer"
	case InProgress:
		return "Em Progresso"
	case Done:
		return "Concluída"
	default:
		return "Desconhecido"
	}
}

// Task é uma struct que agrupa os campos de uma tarefa.
type Task struct {
	Title       string
	Description string
	Status      Status
}

// Stringer para o tipo Task.
// Este metodo é automaticamente chamado para obter a representação em string de uma Task.
func (t Task) String() string {
	return fmt.Sprintf("Título: %s\nDescrição: %s\nStatus: %s", t.Title, t.Description, t.Status)
}

// markAsInProgress marca uma tarefa como "Em Progresso".
func markAsInProgress() {
	listAllTasks() // Mostra as tarefas para o usuário escolher.

	if len(taskList) == 0 {
		fmt.Println("\nNão há tarefas para iniciar.")
		return
	}

	fmt.Print("\nDigite o número da tarefa que deseja INICIAR: ")
	var input string
	fmt.Scanln(&input) // Usa Scanln para capturar a linha inteira, incluindo espaços.

	// Converte a string de entrada para um número inteiro.
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, digite um número.")
		return
	}

	// Ajusta o índice para base 0 (já que o usuário digita a partir de 1).
	idx := index - 1

	// Verifica se o índice é válido.
	if idx < 0 || idx >= len(taskList) {
		fmt.Println("Número de tarefa inválido.")
		return
	}

	// Valida e modifica o status da tarefa.
	if taskList[idx].Status == InProgress {
		fmt.Printf("Tarefa \"%s\" já está EM PROGRESSO.\n", taskList[idx].Title)
	} else if taskList[idx].Status == Done {
		fmt.Printf("Tarefa \"%s\" já está CONCLUÍDA e não pode ser iniciada novamente.\n", taskList[idx].Title)
	} else {
		taskList[idx].Status = InProgress
		fmt.Printf("Tarefa \"%s\" marcada como EM PROGRESSO!\n", taskList[idx].Title)
	}
}

// markAsDone marca uma tarefa como "Concluída".
func markAsDone() {
	listAllTasks() // Mostra as tarefas para o usuário escol1her.

	if len(taskList) == 0 {
		fmt.Println("\nNão há tarefas para marcar como concluídas.") // Mensagem ajustada
		return
	}

	fmt.Print("\nDigite o número da tarefa que deseja CONCLUIR: ")
	var input string
	fmt.Scanln(&input) // Usa Scanln para capturar a linha inteira.

	// Converte a string de entrada para um número inteiro.
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, digite um número.")
		return
	}

	// Ajusta o índice para base 0.
	idx := index - 1

	// Verifica se o índice é válido.
	if idx < 0 || idx >= len(taskList) {
		fmt.Println("Número de tarefa inválido.")
		return
	}

	// Valida e modifica o status da tarefa.
	if taskList[idx].Status == Done {
		fmt.Printf("Tarefa \"%s\" já está CONCLUÍDA.\n", taskList[idx].Title)
	} else if taskList[idx].Status == ToDo {
		fmt.Printf("Tarefa \"%s\" ainda não foi iniciada e não pode ser concluída.\n", taskList[idx].Title)
	} else {
		taskList[idx].Status = Done
		fmt.Printf("Tarefa \"%s\" marcada como CONCLUÍDA!\n", taskList[idx].Title)
	}
}

// listAllTasks lista todas as tarefas no taskList.
func listAllTasks() {
	if len(taskList) == 0 { // Adicionada esta verificação para tarefas vazias
		fmt.Println("\nNenhuma tarefa cadastrada.")
		return
	}
	fmt.Println("\n--- LISTA DE TAREFAS ---")
	for num, task := range taskList {
		fmt.Printf("Tarefa %d:\n", num+1)
		fmt.Println(task)
		fmt.Println("-----------------------")
	}
}

// addNewTask adiciona uma nova tarefa ao taskList.
func addNewTask() {
	var newTask Task

	fmt.Println("\n--- Adicionar Nova Tarefa ---")

	fmt.Print("Digite o TÍTULO da tarefa: ")
	// Usa Scanln para ler a linha inteira, incluindo espaços.
	fmt.Scanln(&newTask.Title)

	fmt.Print("Digite a DESCRIÇÃO da tarefa: ")
	// Usa Scanln para ler a linha inteira, incluindo espaços.
	fmt.Scanln(&newTask.Description)

	newTask.Status = ToDo // Uma nova tarefa começa com status "A Fazer".
	taskList = append(taskList, newTask)
	fmt.Println("Tarefa adicionada com sucesso!")
}

func main() {
	// Loop principal do programa, exibindo o menu e processando as opções.
	for {
		// Loop infinito que será quebrado por 'break' quando a resposta for '0'
		fmt.Println(`
GERENCIAMENTO DE TAREFAS -----------------------------
Digite o número de uma operação:
1 - Adicionar uma tarefa
2 - Iniciar tarefa
3 - Marcar uma tarefa como concluída    
4 - Listar as tarefas.
0 - Encerrar o programa.`)

		var answer string   // Lê a resposta como string para maior robustez
		fmt.Scanln(&answer) // Captura a linha inteira

		trimmedAnswer := strings.TrimSpace(answer) // Remove espaços em branco

		if trimmedAnswer == "0" {
			fmt.Println("Encerrando o programa...")
			break // Sai do loop 'for'
		}

		switch trimmedAnswer {
		case "1":
			addNewTask()
		case "2":
			markAsInProgress()
		case "3":
			markAsDone()
		case "4":
			listAllTasks()
		default:
			fmt.Println("Opção inválida. Por favor, digite um número válido do menu.")
		}
	}
}
