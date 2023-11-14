package main

import "fmt"

func main() {
	// Penerapan Pointer
	var number *int
	var name *string
	fmt.Println(name, number)
	/*
		Ada dua hal penting yang perlu diketahui mengenai pointer:

		Variabel biasa bisa diambil nilai pointernya, caranya dengan menambahkan tanda ampersand (&) tepat sebelum nama variabel. Metode ini disebut dengan referencing.
		Dan sebaliknya, nilai asli variabel pointer juga bisa diambil, dengan cara menambahkan tanda asterisk (*) tepat sebelum nama variabel. Metode ini disebut dengan dereferencing.
	*/
	var numA int = 4
	var numB *int = &numA

	fmt.Println("numA (value)	:", numA)    // 4
	fmt.Println("numA (address)	:", &numA) // 0xc00a

	fmt.Println("numB (value)	:", *numB)  // 4
	fmt.Println("numB (address)	:", numB) // 0xc00a
	/*
		Variabel numberB dideklarasikan bertipe pointer int dengan nilai awal adalah referensi variabel numberA (bisa dilihat pada kode &numberA).
		Dengan ini, variabel numberA dan numberB menampung data dengan referensi alamat memori yang sama.

		Variabel pointer jika di-print akan menghasilkan string alamat memori (dalam notasi heksadesimal), contohnya seperti numberB yang diprint menghasilkan 0xc20800a220.
		Nilai asli sebuah variabel pointer bisa didapatkan dengan cara di-dereference terlebih dahulu (bisa dilihat pada kode *numberB).
	*/

	fmt.Println("")
	fmt.Println("// Efek Perubahan Nilai Pointer")
	// Efek Perubahan Nilai Pointer
	numA = 5

	fmt.Println("numA (value)   :", numA)
	fmt.Println("numA (address) :", &numA)
	fmt.Println("numB (value)   :", *numB)
	fmt.Println("numB (address) :", numB)
	// Variabel numberA dan numberB memiliki referensi memori yang sama. Perubahan pada salah satu nilai variabel tersebut akan memberikan efek pada variabel lainnya.

	fmt.Println("")
	fmt.Println("// Parameter Pointer dgn function")
	// Parameter Pointer dgn function
	var numParam = 4
	fmt.Println("before :", numParam) // 4
	change(&numParam, 10)
	fmt.Println("after :", numParam) // 10

}

func change(original *int, value int) {
	*original = value
}