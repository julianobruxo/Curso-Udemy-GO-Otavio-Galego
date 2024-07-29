package main

import "fmt"

func main() {

	// using var keyword
	// we are not defining
	// any type with variable
	var y = 458

	// taking a pointer variable using
	// var keyword without specifying
	// the type
	var p = &y

	fmt.Println("Value stored in y before changing = ", y)
	fmt.Println("Address of y = ", &y)
	fmt.Println("Value stored in pointer variable p = ", p)
	fmt.Println("Adress of p", &p)

	// this is dereferencing a pointer
	// using * operator before a pointer
	// variable to access the value stored
	// at the variable at which it is pointing
	fmt.Println("Value stored in y(*p) Before Changing = ", *p)

	// changing the value of y by assigning
	// the new value to the pointer
	*p = 500

	fmt.Println("Value stored in y(*p) after Changing = ", y)

}
