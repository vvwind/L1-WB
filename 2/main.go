package main

import (
	"fmt"
	"sync"
)

func main() {
	// Для того, чтобы высчитать квадраты параллельно, а потом вывести, создаем WaitGroup
	var wg sync.WaitGroup

	arr := []int{2, 4, 6, 8, 10}

	resultArr := make([]int, len(arr))

	for i, v := range arr {
		wg.Add(1)
		go func(v int, i int) {
			resultArr[i] = v * v
			defer wg.Done()
		}(v, i)
	}
	wg.Wait()
	// Готово! Теперь выводим результат
	for i, v := range resultArr {
		fmt.Printf("%v в квадрате = %v\n", arr[i], v)
	}
}
