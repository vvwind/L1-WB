package main

import "fmt"

func removeElement(a []int, i int) []int {
	return append(a[:i], a[i+1:]...)
}

func main() {
	arr := []int{0, 1, 2, 3}

	fmt.Println(removeElement(arr, 1))
}
