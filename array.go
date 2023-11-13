package main

import "fmt"

func main() {
	// ARRAY
	var names [4]string
	names[0] = "Muh"
	names[1] = "Ibnu"
	names[2] = "Budi"
	names[3] = "Roma"
	fmt.Println(names[0], names[1], names[2], names[3])
	/*
		Muh Ibnu Budi Roma
	*/
	// Pengisian Elemen Array yang Melebihi Alokasi Awal
	var names2 [2]string
	names2[0] = "Muh"
	names2[1] = "d"
	// names2[2] = "water" // baris kode ini menghasilkan error

	// Inisialisasi Nilai Awal Array
	// cara horizontal
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	// cara vertikal
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	// tanpa jumlah elemen
	var numbers = [...]int{1, 2, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah element \t", len(numbers))

	// Array Multidimensi
	var arrMultidimensi1 = [2][3]int{[3]int{1, 2, 3}, [3]int{3, 2, 1}}

	// boleh tidak dituliskan jumlah datanya.
	var arrMultidimensi2 = [2][3]int{{1, 2, 3}, {3, 2, 1}}
	// [2][3] : 2 adalah jumlah element dan 3 adalah jumlah element di dalam array
	fmt.Println("arrMultidimensi1 \t:", arrMultidimensi1)
	fmt.Println("arrMultidimensi2 \t:", arrMultidimensi2)

	// Perulangan Element Array menggunakan keyword for
	var arrFor = [3]string{"muh", "syahrur", "roma"}
	for i := 0; i < len(arrFor); i++ {
		fmt.Printf("for index %d : element %s\n", i, arrFor[i])
	}
	//for - range
	for i, arrFor := range arrFor {
		fmt.Printf("for - range index %d, elment %s\n", i, arrFor)
	}

	// Penggunaan Variavle Underscore _ Dalam for - range
	// Kegunaan variabel pengangguran, atau underscore (_).
	// Tampung saja nilai yang tidak ingin digunakan ke underscore.
	for _, arrFor := range arrFor {
		fmt.Printf("_underscore elemen : %s\n", arrFor)
	}

	// Alokasi Elemen Array menggunakan keyword make
	var arrMake = make([]string, 2)
	arrMake[0] = "muh"
	arrMake[1] = "syahrur"
	fmt.Println("make arrMake ", arrMake) // [muh syahrur]
	// variabel arrMake tercetak sebagai array string dengan alokasi 2 slot.
}
