package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Assign empty slices
//
//   Assign empty slices to all the slices that you've declared in the previous
//   exercise, and print them here.
//
//
// EXPECTED OUTPUT
//  names    : []string 0 false
//  distances: []int 0 false
//  data     : []uint8 0 false
//  ratios   : []float64 0 false
//  alives   : []bool 0 false
// --------------------------------------
func main() {
	var (
		names     []string
		distances []int
		data      []uint8
		ratios    []float64
		alives    []bool
	)

	names = []string{}
	distances = []int{}
	data = []uint8{}
	ratios = []float64{}
	alives = []bool{}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
}
