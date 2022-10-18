package main

import (
	"fmt"
	"log"

	"github.com/samjin-berkeley/Go-dev/greetings"
)

func main() {
	/*
		Set properties of the predefined Logger, including
		the log entry prefix and a flag to disable printing
		the time, source file, and line number.
	*/

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{
		"Sam",
		"Bob",
		"Younseo",
	}

	messages, err := greetings.Hellos(names)

	// If error was returned, print error message and exit program with status code 1
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map of message sto console
	fmt.Println(messages)
}
