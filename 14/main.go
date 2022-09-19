package main

import "fmt"

func printType(entity interface{}) {
	switch entity.(type) {
	case int:
		fmt.Println("This is a int")
	case float64:
		fmt.Println("This is a float")
	case string:
		fmt.Println("This is a string")
	case bool:
		fmt.Println("This is a bool")
	case chan int:
		fmt.Println("This is a chan int")
	default:
		fmt.Println("This type is not recognized")
	}
}

func main() {
	n := 0
	b := true
	c := 32.4
	str := "hello"
	ch := make(chan int)

	printType(n)
	printType(b)
	printType(c)
	printType(str)
	printType(ch)
	printType(struct{}{})
}
