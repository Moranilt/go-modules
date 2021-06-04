package main

import (
	"fmt"

	greetc "github.com/Moranilt/greetc"
)

func main() {
	msg, err := greetc.Hello("Danil")

	if err != nil {
		fmt.Printf("You have an error! %v", err)
	}

	fmt.Printf(msg)
}
