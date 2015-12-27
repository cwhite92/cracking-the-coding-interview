package main

import "testing"

func TestUnique(t *testing.T) {
	var tests = []struct {
		input  string
		unique bool
	}{
		{"abcde", true},
		{"test", false},
		{"ぁあぃぎじ", true},
		{"せぬがだぬ", false},
		{"The quick brown fox jumps over the lazy dog", false},
		{"", true},
	}

	for _, c := range tests {
		got := unique(c.input)

		if got != c.unique {
			t.Errorf("unique(%q) == %q, want %q", c.input, got, c.unique)
		}
	}
}
