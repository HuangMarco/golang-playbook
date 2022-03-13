package main

import "fmt"

type Creature struct {
	Name     string
	Greeting string
}

func (c Creature) Greet() {
	fmt.Printf("%s says %s", c.Name, c.Greeting)
	fmt.Println("");
}

func main() {
	sammy := Creature{
		Name:     "Sammy",
		Greeting: "Hello!",
	}
	Creature.Greet(sammy)

	// fmt.Println("Hello!")
}
