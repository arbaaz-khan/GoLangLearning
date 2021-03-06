/*
	Singleton class is way to ensure that only one object of a class can be created.
	In go neither we have class nor do we have inbuilt singleton class.
	But we can build one using structs and the properties of a singleton class.
	We have used sync package for it, there is a method called Do available with sync.Once
	with the following definition
	Do (f func)
	It takes a function as an argument and it ensures that that function gets executed only,
	when called for the first time. For all subsequent times that function is not executed.

*/
package main

import (
	"fmt"
	"sync"
)

var instance *singleton
var once sync.Once

type singleton struct {
	name string
}

func GetSingleton() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (s *singleton) getName() string {
	return s.name
}

func main() {
	fmt.Println("Creating object for the first time...")
	sglton1 := GetSingleton()
	fmt.Println("Setting name to arbaaz1 for the first object...")
	sglton1.name = "arbaaz1"
	fmt.Println("Get name from the first object.")
	fmt.Println(sglton1.getName())

	fmt.Println()

	fmt.Println("Trying to create new object...")
	sglton2 := GetSingleton()
	fmt.Println("Checking the value of the member name.")
	fmt.Println(sglton2.getName())
	fmt.Println("Setting name to arbaaz2 to the current object.")
	sglton2.name = "arbaaz2"
	fmt.Println("Get name from the current object.")
	fmt.Println(sglton2.getName())
	fmt.Println("Get name from the first object.")
	fmt.Println(sglton1.getName())

	fmt.Println("\nWe can clearly see that both the objects have the same value, as both are same instances\n To put it in another way only one object has been created.")
}
