package main

import (
	"fmt"
	"sort"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Compare the slices
//
//  1. Split the namesA string and get a slice
//
//  2. Sort all the slices
//
//  3. Compare whether the slices are equal or not
//
//
// EXPECTED OUTPUT
//
//   They are equal.
//
//
// HINTS
//
//   1. strings.Split function splits a string and
//      returns a string slice
//
//   2. Comparing slices: First check whether their length
//      are the same or not; only then compare them.
//
// ---------------------------------------------------------

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesA_slice := strings.Split(namesA, ", ")
	fmt.Printf("%q", namesA_slice)
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	sort.Strings(namesA_slice)
	sort.Strings(namesB)

	if len(namesA_slice) == len(namesB) {
		for i := range namesA_slice {
			if namesA_slice[i] != namesB[i] {
				return
			}
		}
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
