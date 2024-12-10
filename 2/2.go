package main

import "fmt"

func main() {
	a := 5
	b := 10

	a = 10 // Явное переопределение
	b = 5  // Явное переопределение

	fmt.Println(a, b) // Вывод: 10, 5
}
