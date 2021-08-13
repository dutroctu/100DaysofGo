// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func main() {
	var a, b float32
	a = 3.14
	b = 1.12
	fmt.Printf("a:%+v, b: %+v\n", a, b)
	swap(&a, &b)
	fmt.Printf("a:%+v, b: %+v\n", a, b)

	fmt.Printf("a:%p, b: %p\n", &a, &b)

	pa := &a
	pb := &b
	pa, pb = swapAddr(pa, pb)
	fmt.Printf("a:%p, b: %p\n", pa, pb)
}

func swapAddr(pa, pb *float32) (*float32, *float32) {
	return pb, pa
}

func swap(a *float32, b *float32) {
	tmp := *a
	*a = *b
	*b = tmp
}
