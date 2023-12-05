package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello generate a random greeting message for input name
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name is mandatory!")
	}

	message := fmt.Sprintf(randomTemplate(), name)
	return message, nil
}

func randomTemplate() string {

	templates := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	numberOfTemplates := len(templates)
	return templates[rand.Intn(numberOfTemplates)]
}
