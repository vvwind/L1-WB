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

func binSearch(a []int, v int) int {
	r := -1 // not found
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == v {
			r = mid // found
			break
		} else if a[mid] < v {
			start = mid + 1
		} else if a[mid] > v {
			end = mid - 1
		}
	}
	return r
}

func main() {
	slice := []int{50, 100, -100, 0, 29, 85}

	sortedSlice := quicksort(slice)

	fmt.Println(sortedSlice)

	fmt.Println(binSearch(sortedSlice, 0))
	fmt.Println(binSearch(sortedSlice, 100))
	fmt.Println(binSearch(sortedSlice, -100))

}
