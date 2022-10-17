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

/*
Hellos returns a map that associates each of the named people
with a greeting message
params: a slice of strings indicating names
returns: map associating names with a message, error message
*/
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names to messages
	messages := make(map[string]string)

	// Loop through the received slice of names
	// Calling the Hello function to "get" a message for each name
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// Associate message to a name in the map
		messages[name] = message
	}

	return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice (a dynamic array) of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
