package main

import (
	"fmt"
)

type Human struct {
	Name    string
	Surname string
	Age     uint8
}

func (h Human) DoHumanStuff() {
	fmt.Println(`thinks about "Bushido" codex`)
}

func (h Human) Greet() {
	fmt.Printf("Hello, %s %s!\n", h.Name, h.Surname)
}

type Action struct {
	Human
	Profession string
}

func main() {
	a := Action{Human: Human{Name: "Morgan", Surname: "Yu", Age: 30},
		Profession: "scientist"}
	a.Greet()
	fmt.Printf("%s ", a.Name)
	a.DoHumanStuff()
}
