package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	res := 0
	var wg sync.WaitGroup
	wg.Add(len(arr))
	for _, v := range arr {
		go adder(v, &wg, &res)
	}
	wg.Wait()
	fmt.Println("Result is", res)
}

func adder(v int, wg *sync.WaitGroup, res *int) {
	fmt.Println("Adding", v*v)
	*res += v * v
	wg.Done()
}
