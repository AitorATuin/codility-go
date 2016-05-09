package lesson1

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func binaryGap(n int, total int, m int) int {
	if n == 0 {
		return max(total, m)
	}

	if n%2 == 0 {
		fmt.Println("D")
		return binaryGap(n/2, total+1, m)
	}
	return binaryGap(n/2, 0, max(total, m))
}

// BinaryGap within a positive integer N is any maximal sequence of consecutive zeros
// that is surrounded by ones at both ends in the binary representation of N.
// Example: 9 -> 1001 so it contains a binary gap of 2
func BinaryGap(n int) int {
	return binaryGap(n, 0, 0)
}
