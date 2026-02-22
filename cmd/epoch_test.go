package cmd

import "testing"

func TestParseEpoch(t *testing.T) {
	tests := []struct {
		input string
		valid bool
	}{
		{"1700000000", true},
		{"1700000000000", true},
		{"banana", false},
	}
	for _, tt := range tests {
		_, err := parseEpoch(tt.input)
		if tt.valid && err != nil {
			t.Errorf("parseEpoch(%q) unexpected error: %v", tt.input, err)
		}
		if !tt.valid && err == nil {
			t.Errorf("parseEpoch(%q) expected error", tt.input)
		}
	}
}
