package main

import "fmt"

func LucroMaximo(preco []int, i int, comprar bool, memo [][]int) int {
	if i >= len(preco) {
		return 0
	}
	indicecompra := 0
	if comprar {
		indicecompra = 1
	}
	if memo[i][indicecompra] != -1 {
		return memo[i][indicecompra]
	}
	var lucro int
	if comprar {
		// Opção 1: Comprar a ação atual e mudar para modo de venda
		// Opção 2: Não comprar e permanecer no modo de compra
		lucro = max(
			-preco[i]+LucroMaximo(preco, i+1, false, memo), LucroMaximo(preco, i+1, true, memo),
		)
	} else {
		// Opção 1: Vender a ação atual e mudar para modo de compra
		// Opção 2: Não vender e permanecer no modo de venda
		lucro = max(
			preco[i]+LucroMaximo(preco, i+1, true, memo), LucroMaximo(preco, i+1, false, memo),
		)
	}
	memo[i][indicecompra] = lucro
	return lucro
}

// max é uma função auxiliar para obter o máximo entre dois inteiros
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	preco := []int{50, 40, 10, 170, 960, 444, 5, 1000, 532}
	// Inicializamos a tabela de memoização com -1 (valor não calculado)
	memo := make([][]int, len(preco))
	for i := range memo {
		memo[i] = []int{-1, -1} // [0] para venda, [1] para compra
	}

	fmt.Println(LucroMaximo(preco, 0, true, memo))
}
