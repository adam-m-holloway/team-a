package greet

import (
	"errors"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr error
	}{
		{name: "simple name", input: "Ada", want: "Hello, Ada"},
		{name: "name with spaces inside", input: "Ada Lovelace", want: "Hello, Ada Lovelace"},
		{name: "surrounding whitespace trimmed", input: "  Ada  ", want: "Hello, Ada"},
		{name: "unicode name", input: "Zoë", want: "Hello, Zoë"},
		{name: "empty name", input: "", wantErr: ErrEmptyName},
		{name: "whitespace only", input: " \t\n ", wantErr: ErrEmptyName},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Greet(tt.input)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("Greet(%q) error = %v, want %v", tt.input, err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("Greet(%q) unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
