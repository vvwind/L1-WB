package main

import (
	"fmt"
	"strings"
)

func check(str string) bool {
	hash := make(map[string]bool)
	str1 := strings.ToLower(str)
	arr := strings.Split(str1, "")
	fmt.Println(arr)
	for _, v := range arr {
		if _, ok := hash[v]; !ok {
			hash[v] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("abcd: ", check("abcd"))
	fmt.Println("abCdefAaf: ", check("abCdefAaf"))
	fmt.Println("aabcd: ", check("aabcd"))
}
