package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person.
// In Go, a function whose name starts with a capital letter can be called by a function not in the same package
func Hello(name string) (string, error) {
    // Return a greeting that embeds the name in a message.
	if name == ""{
		return "", errors.New("empty name")
	}
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}
func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)
	for _, name := range names{
		message, err := Hello(name)
		if err != nil{
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
func randomFormat() string{

	formats := []string{
		"Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}