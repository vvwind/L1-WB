package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func infiniteWriter(ch chan int, quit chan struct{}) {
	for i := 0; ; i++ {
		select {
		default:
			ch <- i
		case <-quit:
			fmt.Println("Stop writing!")
			return
		}
	}

}

func worker(ch chan int) {
	for {
		message, open := <-ch
		if !open {
			return
		}
		fmt.Println(message)
	}

}

func main() {

	ch := make(chan int)
	quit := make(chan struct{})
	// start writing in channel
	go infiniteWriter(ch, quit)
	// Number of workers from keyboard
	var count int

	_, err := fmt.Fscan(os.Stdin, &count)
	if err != nil {
		log.Panicln(err)
	}
	// making N workers
	for i := 0; i < count; i++ {
		go worker(ch)
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-stop:
		quit <- struct{}{}
		fmt.Println("Stop working!")
	}

}
