package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Get and Set Array Elements
//
//  1. Use the 01-declare-empty exercise
//  2. Remove everything but the array declarations
//
//  3. Assign your best friends' names to the names array
//
//  4. Assign distances to the closest cities to you to the distance array
//
//  5. Assign arbitrary bytes to the data array
//
//  6. Assign a value to the ratios array
//
//  7. Assign true/false values to the alives arrays
//
//  8. Try to assign to the zero array and observe the error
//
//  9. Now use ordinary loop statements for each array and print them
//      (do not use for range)
//
//  10. Now use for range loop statements for each array and print them
//
//  11. Try assigning different types of values to the arrays, break things,
//     and observe the errors
//
//  12. Remove some of the array assignments and observe the loop outputs
//      (zero values)
//
//
// EXPECTED OUTPUT
//
//  Note: The output can change depending on the values that you've assigned to them, of course.
//        You're free to assign any values.
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================

//
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//   FOR RANGES
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================
//
// ---------------------------------------------------------

func main() {
	var (
		names     [3]string
		distances [5]int
		data      [5]uint8
		ratios    [1]float64
		alives    [4]bool
		zero      [0]uint8
	)
	names[0] = "Thanh"
	names[1] = "Loc"
	names[2] = "Vinh"

	distances[0] = 10
	distances[1] = 20
	distances[2] = 30
	distances[3] = 40
	distances[4] = 50

	data[0] = 12
	data[1] = 34
	data[2] = 56
	data[3] = 78
	data[4] = 90

	ratios[0] = 3.14

	alives[0] = true
	alives[1] = false
	alives[2] = true
	alives[3] = false

	fmt.Println("names")
	fmt.Println("====================")
	for i := 0; i < len(names); i++ {
		fmt.Printf("name[%d]: %q\n", i, names[i])
	}
	fmt.Println()

	fmt.Println("distances")
	fmt.Println("====================")
	for i := 0; i < len(distances); i++ {
		fmt.Printf("distances[%d]: %d\n", i, distances[i])
	}
	fmt.Println()

	fmt.Println("data")
	fmt.Println("====================")
	for i := 0; i < len(data); i++ {
		fmt.Printf("data[%d]: %d\n", i, data[i])
	}
	fmt.Println()

	fmt.Println("ratios")
	fmt.Println("====================")
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("ratios[%d]: %.2f\n", i, ratios[i])
	}
	fmt.Println()

	fmt.Println("alives")
	fmt.Println("====================")
	for i := 0; i < len(alives); i++ {
		fmt.Printf("alives[%d]: %t\n", i, alives[i])
	}
	fmt.Println()

	fmt.Println("zero")
	fmt.Println("====================")
	fmt.Printf("%#v", zero)
	fmt.Println()
	fmt.Println()

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
}
