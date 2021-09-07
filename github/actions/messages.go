// Package actions contains utilities for working with github actions
package actions

import (
	"fmt"
)

// Message is a struct used for printing warning/error messages
type Message struct {
	Message string
	File    string
	Line    string
	Col     string
}

// SetOutput can be used to set outputs of your action
func SetOutput(name string, value string) {
	fmt.Printf("::set-output name=%s::%s\n", name, value)
}

// WriteDebug is used to write debug message to the github actions log
func WriteDebug(message string) {
	fmt.Printf("::debug::%s\n", message)
}

// WriteWarning is used to write warning message to the github actions log
func WriteWarning(message Message) {
	message.print("::warning")
}

// WriteError is used to write error message to the github actions log
func WriteError(message Message) {
	message.print("::error")
}
