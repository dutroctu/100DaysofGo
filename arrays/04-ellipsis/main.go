package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Ellipsis
//
//  1. Use the 03-array-literal exercise
//
//  2. Refactor the length of the array literals to ellipsis
//
//    This means: Use the ellipsis instead of defining the array's length
//                manually.
//
// EXPECTED OUTPUT
//   The output should be the same as the 03-array-literal exercise.
// ---------------------------------------------------------

func main() {
	names := [...]string{"Thanh", "Loc", "Vinh"}
	distances := [...]int{10, 20, 30, 40, 50}
	data := [...]uint8{12, 23, 34, 45, 56}
	ratios := [...]float64{3.14}
	alives := [...]bool{true, false, true, false}
	var zero uint8

	fmt.Println("names")
	fmt.Println("====================")
	for i, name := range names {
		fmt.Printf("name[%d]: %q\n", i, name)
	}
	fmt.Println()

	fmt.Println("distances")
	fmt.Println("====================")
	for i, distance := range distances {
		fmt.Printf("distances[%d]: %d\n", i, distance)
	}
	fmt.Println()

	fmt.Println("data")
	fmt.Println("====================")
	for i, d := range data {
		fmt.Printf("data[%d]: %d\n", i, d)
	}
	fmt.Println()

	fmt.Println("ratios")
	fmt.Println("====================")
	for i, ratio := range ratios {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratio)
	}
	fmt.Println()

	fmt.Println("alives")
	fmt.Println("====================")
	for i, alive := range alives {
		fmt.Printf("alives[%d]: %t\n", i, alive)
	}
	fmt.Println()

	fmt.Println("zero")
	fmt.Println("====================")
	fmt.Printf("%#v\n", zero)
	fmt.Println()
}
