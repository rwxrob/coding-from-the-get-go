package main

import "fmt"

func decrement(x int) int      { return x - 1 }
func increment(x int) int      { return x + 1 }
func halve(x int) int          { return x / 2 }
func double(x int) int         { return x * 2 }
func greet(name string) string { return "Hello " + name + "." }

func main() {
	fmt.Println(greet("Doris"))
	fmt.Println(greet("Rob"))
}
