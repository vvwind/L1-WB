package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Реализует постоянную запись данных в канал
func infinitewriter(ch chan int, quit chan struct{}) {
	for i := 0; ; i++ {
		select {
		default:
			ch <- i
		case <-quit:
			return
		}
	}
}

// Выводит произвольные данные
func worker(ch chan int, quit chan struct{}) {
	for {
		select {
		default:
			fmt.Println(<-ch)
		case <-quit:
			return
		}
	}
}

func main() {
	bye := make(chan struct{})
	ch := make(chan int)

	var seconds time.Duration
	fmt.Print("Введите количество секунд: ")
	_, err := fmt.Fscan(os.Stdin, &seconds)
	if err != nil {
		panic(err)
	}

	// Запускаем постоянную запись данных в канал
	go infinitewriter(ch, bye)
	// Запускаем постоянное чтение данных из канала
	go worker(ch, bye)

	// Сделаем так, чтобы процесс не завершался
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	timer := time.NewTimer(time.Second * seconds)

	select {
	case <-stop:
		close(bye)
		fmt.Println("До свидания.")
	case <-timer.C:
		close(bye)
		fmt.Println("До свидания.")
	}
}
