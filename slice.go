// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice
// e os valores dos elementos. ok
// Em seguida, imprima o Slice e a soma dos valores armazenados.
package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Qual o tamanho do slice?")
	fmt.Scan(&tamanho)

	var numeros = make([]int, tamanho)
	fmt.Println("len(numeros)=", len(numeros))

	var add int
	var valor int = 1
	soma := 0

	for valor <= tamanho {
		fmt.Print("Escolha um número.")
		fmt.Scan(&add)
		numeros = append(numeros, add)
		valor++
		soma += add
	}

	fmt.Println(numeros)
	fmt.Println("A soma dos valores indicados é: ", soma)

}
