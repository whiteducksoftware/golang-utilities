/*
Copyright (c) 2021 white duck Gesellschaft für Softwareentwicklung mbH

This code is licensed under MIT license (see LICENSE for details)
*/

// Package actions contains utilities for working with github actions
package actions

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
		output += fmt.Sprintf("line=%s,", message.File)
	}
	if len(message.Col) > 0 {
		output += fmt.Sprintf("col=%s,", message.File)
	}
	fmt.Printf(fmt.Sprintf("%s::%s\n", strings.TrimSuffix(output, ","), message.Message))
}
