package main

import (
	"fmt"
)

func ExampleMain() {

	main()

	// Output:
	// Hello Doris.
	// Hello Rob.
}

func ExampleGreet() {

	fmt.Println(greet("Doris"))
	fmt.Println(greet("Rob"))

	// Output:
	// Hello Doris.
	// Hello Rob.
}

func ExampleDouble() {

	fmt.Println(double(2))
	fmt.Println(double(9))

	// Output:
	// 4
	// 18
}
