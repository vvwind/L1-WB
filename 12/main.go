package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := map[string][]string{}
	for _, v := range arr {
		if _, ok := set[string(v[0])]; !ok {
			set[string(v[0])] = make([]string, 0)
			set[string(v[0])] = append(set[string(v[0])], v)
		} else {
			set[string(v[0])] = append(set[string(v[0])], v)
		}
	}
	fmt.Println(set)
}
