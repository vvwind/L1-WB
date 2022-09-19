package main

import "fmt"

func main() {
	arr1 := []string{"red", "white", "green", "black"}
	arr2 := []string{"orange", "red", "yellow", "white"}

	booleanMap := make(map[string]bool)

	for _, v := range arr1 {
		booleanMap[v] = true
	}

	for _, v := range arr2 {
		if booleanMap[v] == true {
			booleanMap[v] = false
		} else {
			continue
		}
	}
	for k, v := range booleanMap {
		if v == true {
			delete(booleanMap, k)
		}
	}
	for k := range booleanMap {
		fmt.Println(k)
	}

}
