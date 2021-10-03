package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Multiple names
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"1. Hi, %v. Welcome",
		"2. Great to see you %v!",
		"3. Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
