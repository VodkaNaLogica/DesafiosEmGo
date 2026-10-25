package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

var leitor = bufio.NewReader(os.Stdin)

func limparTerminal() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ler(mensagem string) string {
	fmt.Print(mensagem)

	texto, _ := leitor.ReadString('\n')

	return strings.TrimSpace(texto)
}

func menu() {
Loop:
	for {
		println("===== ESTOQUE =====")
		println("\n1. Listar Produtos")
		println("2. Cadastrar Produto")
		println("3. Buscar/Alterar Produto")
		println("4. Sair")
		println("\n===================")
		input := ler("Escolha: ")
		limparTerminal()
		switch input {
		case "1":
			listarProdutos()
		case "2":
			cadastrarProduto()
		case "3":
			buscarProduto()
		case "4":
			break Loop
		}
	}
}

func listarProdutos() {
Loop:
	for {
		fmt.Println("====== ESTOQUE ======")
		fmt.Println("\n1. Retornar")
		if len(ordem) < 1 {
			fmt.Println("\nNão há produtos\ncadastrados ainda")
		} else {
			for _, id := range ordem {
				produto := produtos[id]

				fmt.Printf("\nID: %d\n%s\nR$ %.2f\n%d unidades\n",
					produto.ID,
					produto.Nome,
					produto.Preco,
					produto.Quantidade,
				)
			}
		}
		fmt.Println("\n=====================")
		input := ler("Escolha: ")
		limparTerminal()
		switch input {
		case "1":
			break Loop
		}
	}
}

func cadastrarProduto() {
	defer salvar()
	seletor := 1
	var (
		nome        string
		preco       float32
		quantidade  string
		quantidadeN int
	)

Loop:
	for {
		limparTerminal()
		fmt.Println("= CADASTRAR PRODUTO =")
		fmt.Println("\n1. Retornar")
		switch seletor {
		case 1:
			if nome == "" {
				input := ler("\nEscreva o nome: ")
				if input == "1" {
					limparTerminal()
					break Loop
				}
				nome = input

				limparTerminal()
			} else {
				fmt.Print("\n")
				fmt.Println(nome)
				fmt.Println("\n2. Cancelar")
				fmt.Println("3. Salvar")
				input := ler("\nEscolha: ")
				switch input {
				case "1":
					limparTerminal()
					break Loop
				case "2":
					nome = ""
				case "3":
					seletor++
				}
			}
		case 2:
			if preco == 0 {
				fmt.Print("\n")
				fmt.Println(nome)
				input := ler("\nEscreva o preço: ")

				numero, err := strconv.ParseFloat(input, 32)
				if err != nil {
					fmt.Println("Erro na conversão:", err)
					break
				}

				preco = float32(numero)

				limparTerminal()
			} else {
				fmt.Print("\n")
				fmt.Println(nome)
				fmt.Println(preco)
				fmt.Println("\n2. Cancelar")
				fmt.Println("3. Salvar")
				input := ler("\nEscolha: ")
				switch input {
				case "1":
					limparTerminal()
					break Loop
				case "2":
					preco = 0
				case "3":
					seletor++
				}
			}
		case 3:
			if quantidade == "" {
				fmt.Print("\n")
				fmt.Println(nome)
				fmt.Println(preco)

				input := ler("\nEscreva a quantidade: ")

				numero, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Erro na conversão:", err)
					break
				}
				quantidadeN = numero
				quantidade = input

				limparTerminal()
			} else {
				fmt.Print("\n")
				fmt.Println(nome)
				fmt.Println(preco)
				fmt.Println(quantidade)
				fmt.Println("\n2. Cancelar")
				fmt.Println("3. Salvar")
				input := ler("\nEscolha: ")
				switch input {
				case "1":
					limparTerminal()
					break Loop
				case "2":
					quantidade = ""
				case "3":
					add(nome, preco, quantidadeN)
					limparTerminal()
					break Loop
				}
			}
		}
	}
}

func buscarProduto() {
	var seletor int
	var input string
Loop:
	for {
		seletor = 0
		fmt.Println("====== BUSCAR =======")
		fmt.Println("\n1. Retornar")
		if len(ordem) < 1 {
			fmt.Println("\nNão há produtos\ncadastrados ainda")
		} else {
			fmt.Println("\n2. Buscar por ID")
			fmt.Println("3. Buscar por Nome")
			fmt.Println("4. Buscar por índice")
		}
		fmt.Println("\n=====================")
		input = ler("Escolha: ")
		switch input {
		case "1":
			limparTerminal()
			break Loop
		case "2":
			seletor = 2
		case "3":
			seletor = 3
		case "4":
			seletor = 4
		default:
			limparTerminal()
		}

		if seletor == 2 {
			input := ler("\nEscreva o ID desejado: ")

			numero, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Erro na conversão:", err)
				limparTerminal()
				break
			}

			produto := produtos[numero]
			if produto == nil {
				limparTerminal()
				break
			}

			fmt.Printf("\nID: %d\n%s\nR$ %.2f\n%d unidades\n",
				produto.ID,
				produto.Nome,
				produto.Preco,
				produto.Quantidade,
			)
			editarProduto(produto)
			break Loop
		} else if seletor == 3 {
			input := ler("\nEscreva o nome do produto desejado: ")

			produto := porNome[strings.TrimSpace(strings.ToLower(input))]
			if produto == nil {
				limparTerminal()
				break
			}

			fmt.Printf("\nID: %d\n%s\nR$ %.2f\n%d unidades\n",
				produto.ID,
				produto.Nome,
				produto.Preco,
				produto.Quantidade,
			)
			editarProduto(produto)
			break Loop
		} else if seletor == 4 {
			input := ler("\nEscreva a posição do item desejado: ")

			numero, err := strconv.Atoi(input)
			if err != nil || numero > len(ordem)-1 {
				limparTerminal()
				break
			}

			produto := produtos[ordem[numero]]
			if produto == nil {
				limparTerminal()
				break
			}

			fmt.Printf("\nID: %d\n%s\nR$ %.2f\n%d unidades\n",
				produto.ID,
				produto.Nome,
				produto.Preco,
				produto.Quantidade,
			)
			editarProduto(produto)
			break Loop
		}
	}
}

func editarProduto(produtoEndereco *Produto) {
	defer salvar()
	produto := *produtoEndereco
	fmt.Println("\n2. Editar produto")
	fmt.Println("3. Deletar produto")
	fmt.Println("Pressione enter para sair")
	input := strings.TrimSpace(ler("\nEscolha: "))
	if input == "2" {
		fmt.Printf("\n1. %s\n2. R$ %.2f\n3. %d unidades\n",
			produto.Nome,
			produto.Preco,
			produto.Quantidade,
		)
		campo := ler("\nEscolha para alterar: ")
		switch campo {
		case "1":
			NewName := ler("Qual será o novo nome? ")
			editar(produto.ID, NewName, produto.Preco, produto.Quantidade)
		case "2":
			NewPrice := ler("Qual será o novo preco? ")
			numero, err := strconv.ParseFloat(NewPrice, 32)
			if err != nil {
				limparTerminal()
				break
			}
			editar(produto.ID, produto.Nome, float32(numero), produto.Quantidade)
		case "3":
			NewAmount := ler("Qual será a nova quantidade? ")
			numero, err := strconv.Atoi(NewAmount)
			if err != nil {
				limparTerminal()
				break
			}
			editar(produto.ID, produto.Nome, produto.Preco, numero)
		}
	} else if input == "3" {
		remove(produto.ID)
	}
	limparTerminal()
}

func salvar() {
	dados, err := json.MarshalIndent(produtos, "", "\t")
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	err = os.WriteFile(nomeDoArquivo, dados, 0644)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo JSON:", err)
	}
}

func carregar() {
	_, err := os.Stat(nomeDoArquivo)
	if err == nil {

		dadosLidos, err := os.ReadFile(nomeDoArquivo)
		if err != nil {
			fmt.Println("Erro ao ler o arquivo JSON:", err)
			return
		}
		var produtosCarregados map[int]*Produto

		err = json.Unmarshal(dadosLidos, &produtosCarregados)
		if err != nil {
			fmt.Println("Erro ao carregar o arquivo JSON:", err)
			return
		}

		ordem = []int{}
		porNome = make(map[string]*Produto)

		for id, produto := range produtosCarregados {

			produtos[id] = produto

			ordem = append(ordem, id)

			nomeChave := strings.TrimSpace(strings.ToLower(produto.Nome))
			porNome[nomeChave] = produto
		}

	}
}

func main() {
	defer salvar()
	carregar()
	limparTerminal()
	menu()
}
