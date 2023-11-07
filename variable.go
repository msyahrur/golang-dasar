package main

import "fmt"

func main() {
	//Variables
	//Wajib di deklarasikan

	//Variable declaration
	var msg1 string
	fmt.Println(msg1)
	var msg2 = "Hello, world!"
	fmt.Println(msg2)
	var msg3 string = "Hello, world!"
	fmt.Println(msg3)
	var x1, y1 int
	fmt.Println(x1, y1)
	var x2, y2 int = 1, 2
	fmt.Println(x2, y2)
	var x3, msg4 = 1, "Hello, world!"
	fmt.Println(x3, msg4)
	msg1 = "Hello"
	fmt.Println(msg1)

	//Declaration list
	var (
		x4     int
		y4         = 20
		z4     int = 30
		d4, e4     = 40, "Hello"
		f4, g4 string
	)
	fmt.Println(x4, y4, z4, d4, e4, f4, g4)

}
