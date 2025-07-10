package utils

import (
	"testing"
)

func TestParseTime_Valid(t *testing.T) {
	tests := []struct {
		input    string
		hour     int
		minute  int
	}{
		{"00:00", 0, 0},
		{"09:15", 9, 15},
		{"23:59", 23, 59},
		{"12:30", 12, 30},
	}
	for _, tt := range tests {
		h, m, err := ParseTime(tt.input)
		if err != nil {
			t.Errorf("ParseTime(%q) unexpected error: %v", tt.input, err)
		}
		if h != tt.hour || m != tt.minute {
			t.Errorf("ParseTime(%q) = (%d, %d), want (%d, %d)", tt.input, h, m, tt.hour, tt.minute)
		}
	}
}

func TestParseTime_Invalid(t *testing.T) {
	invalidInputs := []string{
		"24:00",   // invalid hour
		"12:60",   // invalid minute
		"-1:00",   // negative hour
		"12:-1",   // negative minute
		"12",      // missing minute
		"12:30:45",// too many parts
		"abc:def", // non-numeric
		"",        // empty
	}
	for _, input := range invalidInputs {
		_, _, err := ParseTime(input)
		if err == nil {
			t.Errorf("ParseTime(%q) expected error, got nil", input)
		}
	}
}
