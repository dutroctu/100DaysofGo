package main

import "fmt"

type SalaryCalculator interface {
	SalaryCalculator() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) SalaryCalculator() int {
	return p.basicpay + p.pf
}

func (c Contract) SalaryCalculator() int {
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense += v.SalaryCalculator()
	}
	fmt.Printf("Total expense: $%d\n", expense)
}
func main() {
	p1 := Permanent{1, 5000, 20}
	p2 := Permanent{2, 6000, 30}
	c1 := Contract{3, 3000}
	employee := []SalaryCalculator{p1, p2, c1}

	totalExpense(employee)
}
