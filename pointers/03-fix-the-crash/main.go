package main

import "fmt"

type computer struct {
	brand *string
}

func main() {

	c := &computer{}
	change(c, "apple")
	fmt.Printf("brand: %s\n", *c.brand)
}

func change(c *computer, brand string) {
	fmt.Printf("%p\n", c.brand)
	c.brand = &brand
}
