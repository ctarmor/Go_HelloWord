package greetings

import (
	"fmt"
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Gladys")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Failed Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
	fmt.Println("TestHelloName passed")
}

func TestHelloEmptyName(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err != nil {
		t.Fatalf("Failed TestHelloEmptyName")
	}
}
