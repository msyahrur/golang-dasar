package main

import "fmt"

func main() {
	message := greetMe("world")
	fmt.Println(message)
}

func greetMe(name string) string {
	return "Hello, " + name + "!"
}

Shortcut of above (Infers type)
msg := "Hello"
x, msg := 1, "Hello"
Constants
const Phi = 1.618
const Size int64 = 1024
const x, y = 1, 2
const (
	Pi = 3.14
	E  = 2.718
)
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
Constants can be character, string, boolean, or numeric values.
See: Constants https://tour.golang.org/basics/15