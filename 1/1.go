package main

import "fmt"

func main() {
	a := 5
	b := 10

	a = a + b // a = 15
	b = a - b // b = 5
	a = a - b // a = 10

	fmt.Println(a, b) // Вывод: 10, 5
}
