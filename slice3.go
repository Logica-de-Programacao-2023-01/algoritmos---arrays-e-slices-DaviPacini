// Crie um Array de inteiros com 7 elementos.
// Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array.
// Imprima o Array resultante.
package main

import "fmt"

func main() {
	var numeros = []int{1, 4, 7, 13, 22, 46, 69}
	var valor_novo int
	fmt.Println(numeros)

	fmt.Print("Escolha um número para ser adicionado ao array.")
	fmt.Scan(&valor_novo)

	numeros = append([]int{valor_novo}, numeros...)
	fmt.Println(numeros)

	numeros = append(numeros, valor_novo)
	fmt.Println("O array resultante é: ", numeros)
}
