package main

import (
	"fmt"
)

type Human struct {
	moveSpeed int
	firstName string
	lastName  string
}

func (h Human) fullName() string {
	return fmt.Sprintf("Speed of %s %s is %d", h.firstName, h.lastName, h.moveSpeed)
}

type Action struct {
	speedBoost int
	Human
}

func (a Action) Boost(moveSpeed int) string {
	moveSpeed = a.speedBoost * a.Human.moveSpeed
	fmt.Println(a.Human.fullName())
	return fmt.Sprintf("New speed of %s %s is %d", a.Human.firstName, a.Human.lastName, moveSpeed)
}
func main() {
	my_human := Human{
		13,
		"Edgar",
		"Ivanoff",
	}

	my_action := Action{
		4,
		my_human,
	}

	fmt.Println(my_action.Boost(my_human.moveSpeed))
}
