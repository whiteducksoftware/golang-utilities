// Package io contains io utilities for working with github actions
package io

import (
	"fmt"
	"strings"
)

func (message Message) print(prefix string) {
	output := prefix
	if !strings.HasSuffix(output, " ") {
		output += " "
	}

	if len(message.File) > 0 {
		output += fmt.Sprintf("file=%s,", message.File)
	}
	if len(message.Line) > 0 {
		output += fmt.Sprintf("line=%s,", message.Line)
	}
	if len(message.Col) > 0 {
		output += fmt.Sprintf("col=%s,", message.Col)
	}
	fmt.Printf(fmt.Sprintf("%s::%s\n", strings.TrimSuffix(output, ","), message.Message))
}
