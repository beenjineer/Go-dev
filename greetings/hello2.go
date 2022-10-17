package greetings

import (
	"errors"
	"fmt"
)

/*
Hello returns a greeting for the named person.
Hello: Function name, exported
name string: variable_name parameter_type
string, error : return type. Returning multiple variables (https://go.dev/doc/effective_go#multiple-returns)
*/
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
