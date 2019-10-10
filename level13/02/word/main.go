// Package word contains utility functions to count words and their use in a string
package word

import "strings"

// UseCount calculates the number of times each word is used
// returns a mpa where the key is the word and the value is the
// number of times it's been used
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count function counts the number of words in a given string s
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
