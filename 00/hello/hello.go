package main

import (
	"00/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Gladys, from subdirectory")
	fmt.Println(message)
}
