//entry point to program
package main

import "fmt"

func main() {
	n, _ := fmt.Println("Hello", 42, true)
	fmt.Println(n)
	// fmt.Println(err)
	foo()

	fmt.Println("Something more")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()

}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

//control flow
// (1) sequential
// (2) Loop; iterative
// (3) Conditional
