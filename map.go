package main

import "fmt"

func main() {
	/*
		Map adlah tipe data asosiatif yg ada di GO,berbentuk key-value pair.
		Utk setiap data(value) yg disimpan,disiapakan juga keynya.
		key harus uni krn sbg penanda pengaksesan value yg bersangkutan.
	*/
	var chicken map[string]int // Variabel chicken dideklarasikan sebagai map, dengan tipe data key adalah string dan value-nya int

	chicken = map[string]int{} // Kode map[string]int maknanya adalah, tipe data map dengan key bertipe string dan value bertipe int.
	/*
		   	Default nilai variabel map adalah nil.
			Oleh karena itu perlu dilakukan inisialisasi nilai default di awal, caranya cukup dengan tambahkan kurung kurawal pada akhir tipe,
	*/

	// menge-set nilai pada sebuah map
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"]) // januari 50

	fmt.Println("mei", chicken["mei"]) // mei 0
	/*
		belum ada item yang tersimpan menggunakan key "mei".
	*/
	fmt.Println("=============================")

	// Inisialisai Nilai Map
	var data = map[string]int{}
	data["one"] = 1
	fmt.Println("data one: ", data["one"])

	//cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	// cara vertikal
	var chicken2 = map[string]int{
		"jan": 50,
		"feb": 40,
	}
	fmt.Println("chicken1:", chicken1["januari"], chicken2["februari"])
	fmt.Println("chicken2:", chicken1["jan"], chicken2["feb"])

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	// Inisialisasi data menggunakan keyword new, yang dihasilkan adalah data pointer.
	fmt.Println(chicken3, chicken4, chicken5)

	// Iterasi Item Map Menggunakan for - range
	var mapForRange = map[string]int{
		"jan": 40,
		"feb": 30,
		"mar": 23,
	}
	fmt.Println("\n// Iterasi Item Map Menggunakan for - range")
	for key, val := range mapForRange {
		fmt.Println(key, " \t", val)
	}
	// Menghapus Item Map
	fmt.Println("\n// Menghapus Item Map")
	var mapDelete = map[string]int{"jan": 40, "feb": 30}
	fmt.Println(len(mapDelete)) //2
	fmt.Println(mapDelete)

	delete(mapDelete, "jan")
	fmt.Println(len(mapDelete)) //1
	fmt.Println(mapDelete)

	// Deteksi Keberadaan Item dgn Key Tertentu
	var mapCheck = map[string]int{"jan": 40, "feb": 30}
	var valCheck, isExist = mapCheck["mei"]
	// Exist artinya ada
	if isExist {
		fmt.Println(valCheck)
	} else {
		fmt.Println("item is not exists")
	}
	// Kombinasi Slice & Map
	fmt.Println("\n// Kombinasi Slice & Map")
	var MapAndSlice = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, m := range MapAndSlice {
		fmt.Println(m["gender"], m["name"])
	}
}
