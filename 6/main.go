package main

import (
	"fmt"
	"sync"
	"time"
)

func firstRoutine(ch chan string, wg *sync.WaitGroup) {
	for {
		message, ok := <-ch
		if !ok {
			fmt.Println("Shut down!")
			wg.Done()
			return
		}
		fmt.Println(message)
	}
}
func secondRoutine(quit chan struct{}, wg *sync.WaitGroup) {
	n := 1
	for {
		select {
		case <-quit:
			wg.Done()
			fmt.Println("Shut down!")
			return
		default:
			n++
			fmt.Println(n)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan string)
	go firstRoutine(ch, &wg)
	ch <- "Started"
	ch <- "Processing"
	ch <- "First Routine"
	close(ch)
	wg.Add(1)
	quit := make(chan struct{})
	go secondRoutine(quit, &wg)
	time.Sleep(time.Millisecond)
	quit <- struct{}{}
	wg.Wait()

}
