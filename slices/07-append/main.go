package main

import (
	"bytes"
	"fmt"
)

// ---------------------------------------------------------
// EXERCISE: Append
//
//  Please follow the instructions within the code below.
//
// EXPECTED OUTPUT
//  They are equal.
//
// HINTS
//  bytes.Equal function allows you to compare two byte
//  slices easily. Check its documentation: go doc bytes.Equal
// ---------------------------------------------------------

func main() {
	// 1. uncomment the code below
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	for i := range png {
		header = append(header, png[i])
	}
	fmt.Printf("%T %q %q", header, header, png)
	// 2. append elements to header to make it equal with the png slice

	// 3. compare the slices using the bytes.Equal function
	if bytes.Equal(png, header) {
		fmt.Println("They are equal.")
	} else {
		fmt.Println("They are not equal.")
	}

	// 4. print whether they're equal or not
}
