package main

import "fmt"

func main() {
	//Basic types

	//Strings
	strSingle := "Hello"
	strMulti := `Multiline
string`
	fmt.Println(strSingle, strMulti)
	//Strings are of type string.

	//Numbers
	//Typical types
	numInt := 3             // int
	numFloat64 := 3.        // float64
	numComplex128 := 3 + 4i // complex128
	numByte := byte('a')    // byte (alias for uint8)
	fmt.Println(numInt, numFloat64, numComplex128, numByte)

	//Other types
	var uuint uint = 7          // uint (unsigned)
	var pfloat32 float32 = 22.7 // 32-bit float
	fmt.Println(uuint, pfloat32)

	//Arrays
	// var numbers [5]int
	numArrays := [...]int{0, 0, 0, 0, 0}
	fmt.Println(numArrays)
	//Arrays have a fixed size.

	//Slices
	sliceint := []int{2, 3, 4}
	slicebyte := []byte("Hello")
	fmt.Println(sliceint, slicebyte)
	//Slices have a dynamic size, unlike arrays.

	//	Pointers

	b := *getPointer()
	fmt.Println("Value is", b)

	a := new(int)
	*a = 234
	//Pointers point to a memory location of a variable. Go is fully garbage-collected.
	//
	//	See: Pointers https://tour.golang.org/moretypes/1
	//

	//Type conversions
	iConversions := 2
	fConversions := float64(iConversions)
	uConversions := uint(iConversions)
	fmt.Println(iConversions, fConversions, uConversions)
	//See: Type conversions https://tour.golang.org/basics/13
}
func getPointer() (myPointer *int) {
	a := 234
	return &a
}
