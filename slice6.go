//Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e verifique se esse
//valor está presente no Array. Imprima o resultado da busca.

package main

import "fmt"

func main() {
	numeros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var valor int
	fmt.Println("Informe um valor")
	fmt.Scan(&valor)
	for i := 0; i < 10; i++ {
		if numeros[i] == valor {
			fmt.Printf("O valor %d está na array", valor)
			break
		} else {
			continue
		}
	}

}
