package main

import (
	"00/greetings"
	"fmt"
)

type User struct {
	Email          string `yaml:"email"`
	HashedPassword []byte `yaml:"hashed_password"`
	DateofBirth    string `yaml:"dateof_birth"`
	Birthplace     string `yaml:"birthplace"`
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World! 2")
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
