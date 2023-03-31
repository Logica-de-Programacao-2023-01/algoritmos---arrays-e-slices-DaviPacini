// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Print("Escolha um número.")
	fmt.Scan(&x)
	fmt.Print("Escolha um número.")
	fmt.Scan(&y)
	fmt.Print("Escolha um número.")
	fmt.Scan(&z)

	var numeros = [3]int{x, y, z}
	fmt.Println(numeros)
	fmt.Println(x + y + z)
}
