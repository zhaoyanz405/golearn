package main

import (
	"fmt"
	"golearn/greetings"
)

func main() {
	message := greetings.Hello("Glass")
	fmt.Println(message)
}