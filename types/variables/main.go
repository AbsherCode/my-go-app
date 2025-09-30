//package variables
package main

import "fmt"

// use type inference to create a package-level variable "val" and assign it a value of "global"
var val = "global"

// declare package-level variables of type int
var d, e, f int

// declare package-level variables of type bool and override the default values (also known as "zero")
var x, y, z = true, false, true

func main() {
	// print ints
	fmt.Println("d, e, f:", d, e, f)

	// override the default value of a package-level variable
	d = 1_000_000
	fmt.Printf("d: %d\n", d)

	// declare and initialize variables of similar names as booleans but of local scope
	x, y, z := 0, 1.25, "hello"
	fmt.Println("x, y, z:", x, y, z)

	// print the package-level variables x, y, and z
	printVars()

	// create a local variable "val" and assign it an int value
	val := 42

	// print the value of the local variable "val"
	fmt.Printf("%T, local val = %v\n", val, val)

	// print the value of the package-level variable "val"
	printGlobal()

	// update the package-level variable "val" and print it again
	updateGlobal()
	printGlobal()

	// print the pointer address of the local variable "val"
	fmt.Printf("%T, local &val = %v\n", &val, &val)

	// assign a value directly to the pointer address of the local variable "val" and print the value of the local variable "val"
	*(&val) = 100
	fmt.Println("local val = ", val)
}

func printVars() {
	fmt.Println("x, y, z", x, y, z)
}


func updateGlobal() {
	val = "updated global"
}

func printGlobal() {
	fmt.Println("global val = ", val)
}