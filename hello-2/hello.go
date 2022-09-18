package main

import (
	"fmt"
	"app/greetings"
)

func main() {
	message := greetings.Hello("Gladys")
	fmt.Printf(message)
}