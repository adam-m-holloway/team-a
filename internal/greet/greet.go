// Package greet builds greeting messages.
package greet

import (
	"errors"
	"strings"
)

// ErrEmptyName is returned when the name is empty or only whitespace.
var ErrEmptyName = errors.New("name is required")

// Greet returns "Hello, <name>" for a non-blank name.
// Surrounding whitespace is trimmed.
func Greet(name string) (string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return "", ErrEmptyName
	}
	return "Hello, " + trimmed, nil
}
