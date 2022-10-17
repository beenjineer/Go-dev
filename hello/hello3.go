package main

import (
	"fmt"

	"github.com/samjin-berkeley/Go-dev/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
