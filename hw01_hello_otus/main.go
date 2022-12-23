
package main

import "fmt"

func main() {

	// Reversing the string.
	str := "Hello, OTUS!"

	// returns the reversed string.
	strRev := reverse(str)
	fmt.Println(strRev)
}

// function, which takes a string as
// argument and return the reverse of string.
func reverse(str string) (result string) {
	for _, element := range str {
		result = string(element) + result
	}
	return
}
