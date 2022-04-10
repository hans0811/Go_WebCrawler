package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// normal
		{"abcabc", 3},
		{"pwwkew", 3},

		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbb", 1},
		{"abcabcabcd", 4},

		// chinese
		{"一二三一二", 3},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s;"+
				"expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubbstr(b *testing.B) {
	s := "黑化肥發灰會揮發灰化肥揮發會發黑"

	for i := 0; i < 13; i++ {
		s = s + s
	}
	ans := 7

	b.Logf("len(s) = %d", len(s))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s;"+
				"expected %d", actual, s, ans)
		}
	}

}
