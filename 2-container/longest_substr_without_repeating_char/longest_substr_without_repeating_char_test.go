package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		str			string
		expected	int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbb", 1},
		{"abcabcabcd", 4},

		// Chinese support
		{"这里是慕课网", 6},
		{"一二三二一", 3},
	}

	for _, test := range tests {
		actual := getMaxLenOfLongestSubstrWithoutRepeatingChar(test.str)
		if actual != test.expected {
			t.Errorf("got %d for input %str; " + "expected %d",
				actual, test.str, test.expected)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	str := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	for i := 0; i < 13; i++ {
		str = str + str
	}
	b.Logf("len(str) = %d", len(str))
	expected := 8
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := getMaxLenOfLongestSubstrWithoutRepeatingChar(str)
		if actual != expected {
			b.Errorf("got %d for input %str; " + "expected %d",
				actual, str, expected)
		}
	}
}
