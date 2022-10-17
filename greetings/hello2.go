package greetings

import "fmt"

/*
Hello returns a greeting for the named person.
Hello: Function name, exported
name string: variable_name parameter_type
string: return type
*/
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
