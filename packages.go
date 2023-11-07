package main //packages

//Importing
//See: Importing https://tour.golang.org/basics/1
import (
	"fmt" // gives fmt.Println
	"math"
	r "math/rand" // Aliases
)

func main() {
	fmt.Println("packages")

	//import r "math/rand"
	fmt.Println("r ", r.Intn(100))

	//Exporting names
	fmt.Println(math.Pi) // Pi diawali dengan huruf kapital
	//	Exported names begin with capital letters.
	//	See: Exported names https://go.dev/tour/basics/3
}
