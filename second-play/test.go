package main

import (
	"fmt"
)

func main(){
	fmt.Println("input your name!")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Hi %s, I am golang!", name)
}