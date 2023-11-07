package main

import "fmt"

func isTired() bool {
	return true
}
func ifFunc(day string) {
	//Conditional
	if day == "sunday" || day == "saturday" {
		fmt.Println("rest")
	} else if day == "monday" && isTired() {
		fmt.Println("groan")
	} else {
		fmt.Println("work")
	}
	//
	//See: If
	//
	//	Statements in if

	if day := "sunday"; day != "monday" {
		fmt.Println("Uh oh")
	}
	//A condition in an if statement can be preceded with a statement before a ;. Variables declared by the statement are only in scope until the end of the if.
	//
	//See: If with a short statement https://tour.golang.org/flowcontrol/6

	//Switch
	switch day {
	case "sunday":
		// cases don't "fall through" by default!
		//fallthrough
	case "saturday":
		fmt.Println("rest")

	default:
		fmt.Println("work")
	}
	//See: Switch
	//
}
func Loop() {
	//	For loop
	for count := 0; count <= 10; count++ {
		fmt.Println("My counter is at", count)
	}
	//See: For loops https://tour.golang.org/flowcontrol/1
	//
	//	For-Range loop
	entry := []string{"Jack", "John", "Jones"}
	for i, val := range entry {
		fmt.Printf("At position %d, the character %s is present\n", i, val)
	}
	//See: For-Range loops https://gobyexample.com/range

	//	While loop
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	//See: Go's "while" https://tour.golang.org/flowcontrol/3

}
func main() {
	//Flow control
	ifFunc("sunday")
	Loop()

}
