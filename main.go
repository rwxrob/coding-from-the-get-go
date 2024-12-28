package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%v\n", "Hello friend")
	fmt.Printf("%b\n", []byte("Hello friend"))
	fmt.Printf("%U\n", '🌎')
	fmt.Printf("%v\n", string('\U0001F30E'))
	fmt.Printf("%v\n", true)
	fmt.Printf("%v\n", 123)
	fmt.Printf("%v\n", 123.23/234)
}
