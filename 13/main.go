package main

import "fmt"

func main() {

	a := 5
	b := 15
	fmt.Printf("A = %d, B = %d \n", a, b)

	a = a + b // 20
	b = a - b // 5
	a = a - b // 15

	fmt.Printf("A = %d, B = %d", a, b)

}
