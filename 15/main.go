package main

import "fmt"

func createHugeString(amount int) string {
	str := ""
	for i := 0; i < amount; i++ {
		str += "*"
	}
	return str
}

func someFunc() string {

	v := createHugeString(1 << 10)
	// использование глобальной переменной может создать конфликты в многопоточном режиме, возвращаю срез из функции
	return v[:100]
}

func main() {
	// Делаю локальную перменную
	justString := someFunc()
	fmt.Println(justString)
}
