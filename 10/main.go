package main

import "fmt"

func main() {
	degrees := map[int][]float32{}

	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, v := range arr {
		head := int(v / 10)
		if degrees[head] == nil {
			degrees[head] = make([]float32, 0, 8)
			degrees[head] = append(degrees[head], v)
		} else {
			degrees[head] = append(degrees[head], v)
		}

	}
	fmt.Println(degrees)
}
