package main

import "fmt"

func getMultiple() (a string, b string) {
	return "Hello", "World"
}
func Lambdas(x int) {
	//Lambdas
	myfunc := func() bool {
		return x > 10000
	}
	fmt.Println("Lambdas: ", myfunc)
	test1, test2 := func() (string, string) {
		x := []string{"hello", "world"}
		return x[0], x[1]
	}()
	fmt.Println(test1, test2)
	//Functions are first class objects.
}

// Named return values
func NamedReturnValues(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//By defining the return value names in the signature, a return (no args) will return variables with those names.
//
//	See: Named return values https://tour.golang.org/basics/7

func main() {
	//Functions
	Lambdas(1)
	NamedReturnValues(2)

	//Multiple return types
	a, b := getMultiple()
	fmt.Println(a, b)

}
