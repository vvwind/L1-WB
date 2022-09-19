package main

import "fmt"

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			a[0], a[1] = a[1], a[0]
		}
		return a
	}

	left := make([]int, 0)
	right := make([]int, 0)
	half := a[0]

	for _, v := range a[1:] {
		if v < half {
			left = append(left, v)
			continue
		}
		right = append(right, v)
	}

	return append(quicksort(left), append([]int{half}, quicksort(right)...)...)
}

func main() {
	slice := []int{1, 22, -41, 500, -1100, 45, -282}

	fmt.Println(quicksort(slice))
}
