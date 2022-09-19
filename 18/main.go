package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Counter struct {
	i int
	m sync.RWMutex
}

func Run() *Counter {
	return &Counter{
		i: 0,
	}
}

func (v *Counter) Add() {
	v.m.Lock()
	defer v.m.Unlock()
	v.i++
}

func (v *Counter) End() int {
	v.m.RLock()
	defer v.m.RUnlock()
	return v.i
}

func Worker(v *Counter, quit chan struct{}) {
	for {
		select {
		default:
			v.Add()
		case <-quit:
			return
		}
	}
}

func main() {
	inc := Run()
	quit := make(chan struct{})
	for i := 0; i < 10; i++ {
		go Worker(inc, quit)
	}

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-stop:
		quit <- struct{}{}
		fmt.Println("Increments: ", inc.End())
		fmt.Println("Stop working!")

	}
}
