package main

import "testing"

func TestPermutation(t *testing.T) {
	var tests = []struct {
		str1   string
		str2   string
		output bool
	}{
		{"the quick brown fox", "het uqkic wborn ofx", true},
		{"ぁあぃぎじ", "じぎぃあぁ", true},
		{"()[]", "][)(!", false},
		{"ぁあぃぎじ", "じぎぃあぁ!", false},
	}

	for _, test := range tests {
		got := permutation(test.str1, test.str2)

		if got != test.output {
			t.Errorf("permutation(%q, %q) == %t, want %t", test.str1, test.str2, got, test.output)
		}
	}
}
