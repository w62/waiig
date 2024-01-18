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

	s := "the long string"
	i := 10
	fmt.Println("Hello, World!")
	fmt.Println("Hello, World! 2")
	message := greetings.Hello("Gladys")
	fmt.Println(message, " + ", s, " and ", i)
	fmt.Printf("Same thing, but by Printf: %s %d\n", s, i)
}
