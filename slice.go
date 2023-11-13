package main

import "fmt"

func main() {
	/*
		Slice adlah reference elemen array. slice bisa dibut, atau bisa juga dihasilakan dari manipulasi sebuah array ataupun slice lainnya.
		karena mrp data reference,menjadikan perubahan data ditiap elemen slice akan berdampak pada slice lain yang memiliki alamat memori yang sma.

	*/
	// Inisialisasi Slice
	//var slice = []string{"muh", "syahrur", "roma"}
	//fmt.Println(slice) // [muh syahrur roma]
	//fmt.Println(slice[0]) // "muh"

	// Salah satu Perbedaan slice dan array
	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array
	fmt.Println(fruitsA, fruitsB, fruitsC)
	/*
		Hubungan Slice Dengan Array & Operasi Slice
			Sebenarnya slice dan array tidak bisa dibedakan karena merupakan sebuah kesatuan.
			Array adalah kumpulan nilai atau elemen, sedang slice adalah referensi tiap elemen tersebut.
			Slice bisa dibentuk dari array yang sudah didefinisikan, caranya dengan memanfaatkan teknik 2 index untuk mengambil elemen-nya.
			Contoh bisa dilihat pada kode berikut.
	*/
	var fruits = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits[0:2]
	fmt.Println(newFruits) // ["apple", "grape"]
	/*
		Kode fruits[0:2] maksudnya adalah pengaksesan elemen dalam slice fruits yang dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2.
		Elemen yang memenuhi kriteria tersebut akan didapat, untuk kemudian disimpan pada variabel lain sebagai slice baru.
		Pada contoh di atas, newFruits adalah slice baru yang tercetak dari slice fruits, dengan isi 2 elemen, yaitu "apple" dan "grape"

		Kode		Output							Penjelasan
		fruits[0:2]	[apple, grape]					semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
		fruits[0:4]	[apple, grape, banana, melon]	semua elemen mulai indeks ke-0, hingga sebelum indeks ke-4
		fruits[0:0]	[]								menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
		fruits[4:4]	[]								menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
		fruits[4:0]	[]								error, pada penulisan fruits[a:b] nilai a harus lebih kecil atau sama dengan b
		fruits[:]	[apple, grape, banana, melon]	semua elemen
		fruits[2:]	[banana, melon]					semua elemen mulai indeks ke-2
		fruits[:2]	[apple, grape]					semua elemen hingga sebelum indeks ke-2
	*/
	// Slice Merupakan Tipe Data Reference
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]
	// pada output di atas, elemen yang sebelumnya bernilai "grape" pada variabel fruits, aFruits, bFruits, aaFruits, dan baFruits;
	// Seluruhnya berubah menjadi "pinnaple", karena memiliki referensi yang sama.

	// Fungsi len() digunakan untuk menghitung jumlah elemen slice yang ada
	fmt.Println("len: ", len(fruits)) //4

	//Fungsi cap()  digunakan untuk menghitung lebar atau kapasitas maksimum slice.
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits1 = fruits[0:3]
	fmt.Println(len(aFruits1)) // len: 3
	fmt.Println(cap(aFruits1)) // cap: 4

	var bFruits1 = fruits[1:4]
	fmt.Println(len(bFruits1)) // len: 3
	fmt.Println(cap(bFruits1)) // cap: 3
	/*
		Kode			Output					len()	cap()
		fruits[0:4]		[buah buah buah buah]	4		4
		aFruits1[0:3]	[buah buah buah ----]	3		4
		bFruits1[1:4]	---- [buah buah buah]	3		3
	*/

	// Fungsi append() digunakan untuk menambahkan elemen pada slice.
	var appendFruits = []string{"apple", "grape", "banana"}
	var appendFruits2 = append(fruits, "papaya")

	fmt.Println(appendFruits)  // ["apple", "grape", "banana"]
	fmt.Println(appendFruits2) // ["apple", "grape", "banana", "papaya"]
	/*
		Hal yang perlu diketahui dalam penggunaan fungsi ini.
		 - Ketika jumlah elemen dan lebar kapasitas adalah sama (len(fruits) == cap(fruits)), maka elemen baru hasil append() merupakan referensi baru.
		 - Ketika jumlah elemen lebih kecil dibanding kapasitas (len(fruits) < cap(fruits)), elemen baru tersebut ditempatkan ke dalam cakupan kapasitas, menjadikan semua elemen slice lain yang referensi-nya sama akan berubah nilainya.

	*/
	var _ = /*appendFruits*/ []string{"apple", "grape", "banana"}
	var appendbFruits = appendFruits[0:2]

	fmt.Println(cap(appendbFruits)) // 3
	fmt.Println(len(appendbFruits)) // 2

	fmt.Println(appendFruits)  // ["apple", "grape", "banana"]
	fmt.Println(appendbFruits) // ["apple", "grape"]

	var appendcFruits = append(appendbFruits, "papaya")

	fmt.Println(appendFruits)  // ["apple", "grape", "papaya"]
	fmt.Println(appendbFruits) // ["apple", "grape"]
	fmt.Println(appendcFruits) // ["apple", "grape", "papaya"]

	// Fungsi copy() digunakan untuk men-copy elements slice pada src (parameter ke-2), ke dst (parameter pertama).
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pinnaple apple
	fmt.Println(src) // watermelon pinnaple apple orange
	fmt.Println(n)   // 3
	//== contoh lain
	dst2 := []string{"potato", "potato", "potato"}
	src2 := []string{"watermelon", "pinnaple"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2) // watermelon pinnaple potato
	fmt.Println(src2) // watermelon pinnaple
	fmt.Println(n2)   // 2

	/*
		Pengaksesan Elemen Slice Dengan 3 Indeks
		3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya.
		Cara menggunakannnya yaitu dengan menyisipkan angka kapasitas di belakang, seperti fruits[0:1:1].
		Angka kapasitas yang diisikan tidak boleh melebihi kapasitas slice yang akan di slicing.
	*/

	var fruits3Idx = []string{"apple", "grape", "banana"}
	var aFruits3Idx = fruits3Idx[0:2]
	var bFruits3Idx = fruits3Idx[0:2:2]

	fmt.Println(fruits3Idx)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits3Idx)) // len: 3
	fmt.Println(cap(fruits3Idx)) // cap: 3

	fmt.Println(aFruits3Idx)      // ["apple", "grape"]
	fmt.Println(len(aFruits3Idx)) // len: 2
	fmt.Println(cap(aFruits3Idx)) // cap: 3

	fmt.Println(bFruits3Idx)      // ["apple", "grape"]
	fmt.Println(len(bFruits3Idx)) // len: 2
	fmt.Println(cap(bFruits3Idx)) // cap: 2

}
