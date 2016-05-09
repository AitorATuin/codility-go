package lesson1

import "testing"

func TestBinaryGap(t *testing.T) {
	cases := []struct {
		in, expected int
	}{
		{1, 0},
		{2, 1},
		{9, 2},
		{16, 4},
	}
	for _, c := range cases {
		got := BinaryGap(c.in)
		if got != c.expected {
			t.Errorf("BinaryGap(%d) == %d, but expected %d", c.in, got, c.expected)
		}
	}
}
