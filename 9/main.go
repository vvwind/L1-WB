package main

import (
	"fmt"
	"time"
)

func conv(ch chan int, ch_mp chan int, quit chan struct{}) {

	go func() {
		for {
			select {
			case data := <-ch:
				ch_mp <- data * data
			case <-quit:
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case data := <-ch_mp:

				fmt.Println(data)
			case <-quit:
				return
			}
		}
	}()
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int)
	ch_mp := make(chan int)
	quit := make(chan struct{})
	conv(ch, ch_mp, quit)
	for _, v := range arr {
		go func(v int) {
			ch <- v
		}(v)
	}

	time.Sleep(time.Second * 1)
	quit <- struct{}{}
}
