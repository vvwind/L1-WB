package main

import (
	"fmt"
	"log"
)

func main() {
	var number_for_xor int64
	var number int64
	var ind uint64
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Panicln("error scanning!")
	}
	_, err2 := fmt.Scan(&ind)
	if err2 != nil {
		log.Panicln("error scanning!")
	}

	number_for_xor = (1 << (ind - 1))

	output := number_for_xor ^ number

	fmt.Printf("%b\n", number)
	fmt.Printf("%b\n", number_for_xor)
	fmt.Printf("%b\n", output)
}
