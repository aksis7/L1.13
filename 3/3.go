package main

import "fmt"

func main() {
	a := 5
	b := 10

	a, b = b, a // Одновременный обмен значений

	fmt.Println(a, b) // Вывод: 10, 5
}
