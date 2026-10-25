package main

import "fmt"

type Produto struct {
	ID         int
	Nome       string
	Preco      float32
	Quantidade int
}

var produtos = make(map[int]*Produto)
var porNome = make(map[string]*Produto)
var ordem []int

var proximoId int = 1

var nomeDoArquivo string = "produtos.json"

func add(nome string, preco float32, quantidade int) {

	if _, existe := porNome[nome]; existe {
		fmt.Println("Já existe um produto com esse nome.")
		return
	}

	produto := &Produto{
		ID:         proximoId,
		Nome:       nome,
		Preco:      preco,
		Quantidade: quantidade,
	}

	produtos[produto.ID] = produto
	porNome[produto.Nome] = produto
	ordem = append(ordem, produto.ID)

	proximoId++
}

func remove(id int) {
	produto, existe := produtos[id]

	if !existe {
		fmt.Println("Produto não encontrado")
		return
	}

	delete(produtos, id)
	delete(porNome, produto.Nome)

	for i, idOrdem := range ordem {
		if idOrdem == id {
			ordem = append(ordem[:i],
				ordem[i+1:]...)
			break
		}
	}

	fmt.Println("Produto removido!")
}

func editar(id int, novoNome string, novoPreco float32, novaQuantidade int) {
	produto, existe := produtos[id]
	if !existe {
		fmt.Println("Produto não existe.")
		return
	}

	produto.Nome = novoNome
	produto.Preco = novoPreco
	produto.Quantidade = novaQuantidade
}
