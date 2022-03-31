package main

import "fmt"
import . "arithmetic/base/day05"

func main() {
	s := "abab"
	p := "aab"
	fmt.Println(FindAnagrams(s, p))
}
