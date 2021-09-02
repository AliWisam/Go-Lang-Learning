package main

import "fmt"

var y = 42
var z = "any string value"
var a = `james said, 

"Shaken, 

not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
