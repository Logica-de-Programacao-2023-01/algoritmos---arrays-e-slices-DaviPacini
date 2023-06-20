//Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas. Solicite ao usuário que informe os valores de
//cada elemento da matriz. Em seguida, imprima a matriz resultante.

package main

import "fmt"

func main() {
	var (
		matriz [3][2]int
		n1     int
		n2     int
		l      = 0
		c      = 0
	)
	for i := 0; i < 3; i++ {
		fmt.Println("Informe os valores da primeira linha")
		fmt.Scan(&n1)
		fmt.Scan(&n2)
		matriz[l][c] = n1
		matriz[l][c+1] = n2
		l += 1
	}
	fmt.Print(matriz)
}
