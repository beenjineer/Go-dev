package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/*
Hello returns a greeting for the named person. Notice that the function starts with a capital
letter. This function can only be called by a external function not in the same package.
Hello: Function name, exported
name string: variable_name parameter_type
string, error : return type. Returning multiple variables (https://go.dev/doc/effective_go#multiple-returns)
*/
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message. (https://pkg.go.dev/fmt#Sprintf)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
