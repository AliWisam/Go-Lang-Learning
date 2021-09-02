package main

import "fmt"

var y string
var z int

func main() {
	fmt.Println("y", y)
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"
	fmt.Println("Printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println("z", z)
	fmt.Printf("%T\n", z)

	z = 73
	fmt.Println("Printing the value of z", z, "ending")
	fmt.Printf("%T\n", z)
}
