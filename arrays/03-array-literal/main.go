package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Refactor to Array Literals
//
//  1. Use the 02-get-set-arrays exercise
//
//  2. Refactor the array assignments to array literals
//
//    1. You would need to change the array declarations to array literals
//
//    2. Then, you would need to move the right-hand side of the assignments,
//       into the array literals.
//
// EXPECTED OUTPUT
//   The output should be the same as the 02-get-set-arrays exercise.
// ---------------------------------------------------------

func main() {
	names := [3]string{"Thanh", "Loc", "Vinh"}
	distances := [5]int{10, 20, 30, 40, 50}
	data := [5]uint8{12, 23, 34, 45, 56}
	ratios := [1]float64{3.14}
	alives := [4]bool{true, false, true, false}
	var zero [0]uint8

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
