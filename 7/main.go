package main

import (
	"fmt"
	"sync"
)

func writer(my_map map[string]int, wg *sync.WaitGroup, val string, number int, mu *sync.RWMutex) {
	mu.Lock()
	my_map[val] = number
	defer mu.Unlock()
	defer wg.Done()
}

func main() {
	my_map := make(map[string]int)
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(3)
	items := []string{"Hello", "World", "!!!"}
	for i, v := range items {
		go writer(my_map, &wg, v, i, &mu)
	}

	wg.Wait()
	for k, v := range my_map {
		fmt.Println(k, v)
	}

}
