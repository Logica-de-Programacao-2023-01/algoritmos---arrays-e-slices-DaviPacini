//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro.
//Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.

package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	var n int
	fmt.Println("Informe um número")
	fmt.Scan(&n)
	for i := 0; i < len(numeros); i++ {
		if n != numeros[i] {
			numeros = append(numeros, n)
			break
		} else {
			fmt.Println("O número ja ")
			continue
		}

	}
	fmt.Print(numeros)
}
