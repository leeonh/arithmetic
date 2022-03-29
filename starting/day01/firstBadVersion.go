package main

import "fmt"

func firstBadVersion(n int) int {
	l, r := 0, n
	for l < r {
		m := (l + r) >> 1
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
func isBadVersion(n int) bool {
	if n == 3 {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(firstBadVersion(5))
}
