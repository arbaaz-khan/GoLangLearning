/* *****************************************************************************
 *  Name:      Arbaaz Khan
 *  Language:  go
 *
 *  Description:  In this program we have explored functions in go.
 *
 *  % go run functions.go
 *
***************************************************************************** */
package main

import (
	"fmt"
)

func main() {
	func1()
	println(func2())
	r1, r2 := func3()
	println("r1=", r1, "r2=", r2)
	r3, r4, r5 := func4() // Or we can also ignore certain number of returned values by using _ (blank), ex, r3, r4, _ := func4()
	println("area=", r3, "x=", r4, "y=", r5)
	func5(5, true)
	s, p := func6(1, 2, 3)
	fmt.Println("SUM", s, " PRODUCT:", p)
	fmt.Println("Methods")
	person := new(Person)
	// person := Person{}
	fmt.Println(person.tellYourName())
}

// Function that doesn't return anything
func func1() {
	println("I am in func1")
}

// Function that returns a single value
func func2() string {
	result := "I am in func2"
	return result
}

// Function that returns multiple values
func func3() (string, int) {
	result := "I am in func3"
	x := 5
	return result, x
}

// Function with named return values
func func4() (area, x int, y int) {
	x = 5
	y = 2
	area = x * y
	return // No need to mention what to return here as we have already mentioned with name in the
}

// Function with arguments
func func5(x int, y bool) {
	println("x=", x, "y=", y)
}

// Variadic function
// Multiple arguments of same type can be passed
// Notice how multiple values are returned
func func6(nums ...int) (int, int) { //nums is a slice of int
	fmt.Print(nums, " ")
	sum := 0
	product := 1
	for _, num := range nums {
		sum += num
		product *= num
	}
	return sum, product
}

// Methods
// Methods are like functions but they are bound to structs.
type Person struct {
	name string
}

func (p Person) tellYourName() string {
	return p.name
}
