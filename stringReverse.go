// Reverse a string in go without using the reverse function
package main

import (
	"strings"
)

func main() {
	str := "Vedant Madane"
	rev := ""
	for i := len(str) - 1; i >= 0; i-- {
		rev += strings.TrimSpace(string(str[i]))
		if i == 0 {
			println(rev)
		}

	}
}
