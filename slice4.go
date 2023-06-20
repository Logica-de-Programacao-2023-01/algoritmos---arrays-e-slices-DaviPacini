//Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos.
//Em seguida, imprima o Slice e a soma dos valores armazenados.

package main

import "fmt"

func main() {
	var (
		n     int
		valor int
	)
	fmt.Println("Informe a quantidade de elementos que serão dados:")
	fmt.Scan(&n)
	if n > 0 {
		numeros := []int{}
		for i := 0; i < n; i++ {
			fmt.Println("Informe o valor:")
			fmt.Scan(&valor)
			numeros = append(numeros, valor)
		}
		fmt.Print(numeros)
	} else {
		fmt.Println("O valor dado é invalido")
	}
}
