package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	var name string = "Gladys"
	// regexp \b means boundary (https://javascript.info/regexp-boundary#:~:text=When%20the%20regexp%20engine%20(program,is%20a%20word%20character%20%5Cw%20.)
	want := regexp.MustCompile(`\b` + name + `\b`)

	msg, err := Hello("Gladys")

	if !want.MatchString(msg) || err != nil {
		// %q prints the value as a quoted string. %#q prints the value as a quoted string as Go syntax.
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with empty name, checking
// for a valid return value which is an error
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want match for "", error`, msg, err)
	}
}
