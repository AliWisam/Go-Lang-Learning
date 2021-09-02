package main

import (
	"fmt"
)

var y = 43

//declare variable z of type int and assigns 0 value of Type int to z
var z int

func main() {
	x := 42
	fmt.Println(x)

	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("this is var value: ", y)
}
