package main

import (
	"fmt"
)

type computer struct {
	brand string
}

func change(comp *computer, brand string) {
	comp.brand = brand
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var comp *computer
	// compare the pointer variable to a nil value, and say it's nil
	if comp == nil {
		fmt.Println("Comp is nil")
	}
	// create an apple computer by putting its address to a pointer variable
	apple := &computer{brand: "Apple"}
	// put the apple into a new pointer variable
	newApple := apple
	// compare the apples: if they are equal say so and print their addresses
	if newApple == apple {
		fmt.Printf("apple %p equal newApple %p\n", apple, newApple)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony := &computer{brand: "Sony"}
	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if sony != apple {
		fmt.Printf("apple %p is not equal sony %p\n", apple, sony)
	}
	// put apple's value into a new ordinary variable
	appleValue := *apple

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("apple: %p %p\n", &apple, apple)
	fmt.Printf("appleValue: %p\n", &appleValue)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if appleValue == *apple {
		fmt.Println("apple equal to appleValue")
		fmt.Printf("apple %+v, appleValue: %+v\n", *apple, appleValue)
	}
	// print the values:
	// the value that is pointed by the apple and the new variable

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sony, "hp")
	fmt.Printf("sony:  %s\n", sony.brand)
	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	fmt.Printf("apple: %+v\n", valueOf(apple))
	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("dell's address            : %p\n", newComputer("dell"))
	fmt.Printf("lenovo's address          : %p\n", newComputer("lenovo"))
	fmt.Printf("acer's address            : %p\n", newComputer("acer"))

}

func valueOf(c *computer) computer {
	return *c
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}
