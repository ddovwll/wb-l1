package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (human Human) HumanInfo() {
	fmt.Printf("%s %s: %v years old\n", human.Name, human.Surname, human.Age)
}

func (human Human) SaySomething(phrase string) {
	fmt.Printf("%s %s said: %s\n", human.Name, human.Surname, phrase)
}

type Action struct {
	Human
}

func main() {
	human := Human{
		Name:    "Robin",
		Surname: "Bobin",
		Age:     18,
	}

	human.HumanInfo()
	human.SaySomething("Hi!")

	action := Action{
		Human{
			Name:    "Walter",
			Surname: "White",
			Age:     52,
		},
	}

	action.HumanInfo()
	action.SaySomething("Say my name.")
}
