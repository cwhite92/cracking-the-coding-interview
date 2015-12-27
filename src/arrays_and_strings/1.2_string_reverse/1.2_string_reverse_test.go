package main

import "testing"

func TestReverse(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"abcde", "edcba"},
		{"hello world", "dlrow olleh"},
		{"ぁあぃぎじ", "じぎぃあぁ"},
		{"", ""},
		{"()[]", "][)("},
	}

	for _, c := range tests {
		got := reverse(c.input)

		if got != c.output {
			t.Errorf("reverse(%q) == %q, want %q", c.input, got, c.output)
		}
	}
}
