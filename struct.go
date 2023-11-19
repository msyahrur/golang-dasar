package main

import "fmt"

/*
Go tidak memiliki class yang ada di bahasa-bahasa strict OOP lain. Tapi Go memiliki tipe data struktur yang disebut dengan Struct.

Struct adalah kumpulan definisi variabel (atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data baru dengan nama tertentu. Property dalam struct, tipe datanya bisa bervariasi. Mirip seperti map, hanya saja key-nya sudah didefinisikan di awal, dan tipe data tiap itemnya bisa berbeda.

Dari sebuah struct, kita bisa buat variabel baru, yang memiliki atribut sesuai skema struct tersebut. Kita sepakati dalam buku ini, variabel tersebut dipanggil dengan istilah object atau object struct.

Konsep struct di golang mirip dengan konsep class pada OOP, meski sebenarnya berbeda. Di sini penulis menggunakan konsep OOP sebagai analogi, dengan tujuan untuk mempermudah dalam mencerna isi chapter ini.

Dengan memanfaatkan struct, grouping data akan lebih mudah, selain itu rapi dan gampang untuk di-maintain.
*/

// Deklarasi Struct
// Struct student dideklarasikan memiliki 2 property, yaitu name dan grade. Objek yang dibuat dengan struct ini nantinya memiliki skema atau struktur yang sama.
type student struct {
	name  string
	grade int
}

func main() {
	// Penerapan Struct
	var s1 student
	s1.name = "muh"
	s1.grade = 2
	fmt.Println("name :", s1.name)
	fmt.Println("name", s1.grade)

	fmt.Println("\n// Inisialisasi Object Struct")
	// Inisialisasi Object Struct
	var objStruct1 = student{}
	objStruct1.name = "muh"
	objStruct1.grade = 2

	var objStruct2 = student{"syahrur", 3}
	var objStruct3 = student{name: "roma"}
	var objStruct4 = student{name: "wayne", grade: 2}
	var objStruct5 = student{grade: 2, name: "bruce"}
	fmt.Println("student 1:", objStruct1.name)
	fmt.Println("student 2:", objStruct2.name)
	fmt.Println("student 3:", objStruct3.name)
	fmt.Println("student 4:", objStruct4.name)
	fmt.Println("student 5:", objStruct5.name)
	/*
		variabel objStruct1 menampung objek cetakan student. Variabel tersebut kemudian di-set nilai property-nya.
		Variabel objStruct2 dideklarasikan dengan metode yang sama dengan objStruct1, pembedanya di objStruct2 nilai propertinya di isi langsung ketika deklarasi.
		deklarasi objStruct3, dilakukan juga pengisian property ketika pencetakan objek. Hanya saja, yang diisi hanya name saja. Cara ini cukup efektif jika digunakan untuk membuat objek baru yang nilai property-nya tidak semua harus disiapkan di awal. Keistimewaan lain menggunakan cara ini adalah penentuan nilai property bisa dilakukan dengan tidak berurutan.
	*/

	fmt.Println("\n// Variable Object Pointer")
	// Variable Object Pointer
	//		Objek yang dibuat dari tipe struct bisa diambil nilai pointer-nya, dan bisa disimpan pada variabel objek yang bertipe struct pointer.
	var objPtr1 = student{name: "muh", grade: 2}
	var objPtr2 *student = &objPtr1
	fmt.Println("student 1, name :", objPtr1.name)
	fmt.Println("student 2, name :", objPtr2.name)

	objPtr2.name = "syahrur"
	fmt.Println("student 1, name :", objPtr1.name)
	fmt.Println("student 2, name :", objPtr2.name)

	/*
		objPtr2 adalah variabel pointer hasil cetakan struct student. objPtr2 menampung nilai referensi objPtr1, menjadikan setiap perubahan pada property variabel tersebut, akan juga berpengaruh pada variabel objek objPtr1.
		objPtr2 bukan variabel asli, property nya tetap bisa diakses seperti biasa. Inilah keistimewaan property dalam objek pointer, tanpa perlu di-dereferensi nilai asli property tetap bisa diakses. Pengisian nilai pada property tersebut juga bisa langsung menggunakan nilai asli, contohnya seperti objPtr2.name = "syahrur"
	*/

	fmt.Println("\n// Embedded Struct")
	// Embedded Struct
	var embedS1 = student2{}
	embedS1.name = "syahrur"
	embedS1.age = 21
	embedS1.grade = 2

	fmt.Println("name  :", embedS1.name)  // syahrur
	fmt.Println("age   :", embedS1.age)   // 21
	fmt.Println("grade :", embedS1.grade) // 21

	/*
		Embedded struct adalah mutable, nilai property-nya nya bisa diubah.

		Khusus untuk properti yang bukan properti asli (properti turunan dari struct lain),
		bisa diakses dengan cara mengakses struct parent-nya terlebih dahulu, contohnya s1.person.age.
		Nilai yang dikembalikan memiliki referensi yang sama dengan s1.age.
	*/
	fmt.Println("age   :", embedS1.person.age)

	fmt.Println("\n// Embedded Struct Dengan Nama Property Yang Sama")
	// Embedded Struct Dengan Nama Property Yang Sama
	var embedS2 = student3{}
	embedS2.name = "wick"
	embedS2.age = 21         // age of student
	embedS2.person3.age = 22 // age of person

	fmt.Println(embedS2.name)        // wick
	fmt.Println(embedS2.age)         // 21
	fmt.Println(embedS2.person3.age) // 22

	// Pengisian Nilai Sub-Struct
	// Pengisian nilai property sub-struct bisa dilakukan dengan langsung memasukkan variabel objek yang tercetak dari struct yang sama.
	fmt.Println("\n// Pengisian Nilai Sub-Struct")
	var pSubStruct = person{name: "wick", age: 21}
	var sSubStruct = student2{person: pSubStruct, grade: 2} // pSubStruct = Pengisian Nilai Sub-Struct

	fmt.Println("name  :", sSubStruct.name)
	fmt.Println("age   :", sSubStruct.age)
	fmt.Println("grade :", sSubStruct.grade)
	//Pada deklarasi s1, property person diisi variabel objek p1.

	/*
	   A.24.8. Anonymous Struct
	   Anonymous struct adalah struct yang tidak dideklarasikan di awal sebagai tipe data baru, melainkan langsung ketika pembuatan objek. Teknik ini cukup efisien untuk pembuatan variabel objek yang struct-nya hanya dipakai sekali.

	*/

	fmt.Println("\n// Anonymous Struct")
	// anonymous struct tanpa pengisian property
	var sAnony1 = struct {
		personAnonymous
		grade int
	}{}
	sAnony1.personAnonymous = personAnonymous{"wick", 21}
	sAnony1.grade = 2

	// anonymous struct dengan pengisian property
	var sAnony2 = struct {
		personAnonymous
		grade int
	}{
		personAnonymous: personAnonymous{"syahrur", 22},
		grade:           2,
	}

	fmt.Println("name  :", sAnony1.personAnonymous.name)
	fmt.Println("age   :", sAnony1.personAnonymous.age)
	fmt.Println("grade :", sAnony1.grade)
	fmt.Println("sAnony2  :", sAnony2.personAnonymous.name)
	/*
	   Pada kode di atas, variabel sAnony1 langsung diisi objek anonymous struct yang memiliki property grade, dan property person yang merupakan embedded struct.

	   Salah satu aturan yang perlu diingat dalam pembuatan anonymous struct adalah, deklarasi harus diikuti dengan inisialisasi. Bisa dilihat pada sAnony1 setelah deklarasi struktur struct, terdapat kurung kurawal untuk inisialisasi objek. Meskipun nilai tidak diisikan di awal, kurung kurawal tetap harus ditulis.
	*/

	/*
		A.24.9. Kombinasi Slice & Struct
		Slice dan struct bisa dikombinasikan seperti pada slice dan map, caranya penggunaannya-pun mirip, cukup tambahkan tanda [] sebelum tipe data pada saat deklarasi.
	*/
	fmt.Println("\n// Kombinasi Slice & Struct")
	type pKombinasiSliceStruct struct {
		name string
		age  int
	}

	var allStudents = []pKombinasiSliceStruct{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	/*
		A.24.10. Inisialisasi Slice Anonymous Struct
		Anonymous struct bisa dijadikan sebagai tipe sebuah slice. Dan nilai awalnya juga bisa diinisialisasi langsung pada saat deklarasi. Berikut adalah contohnya:
	*/
	var sSliceAnonymousStruct = []struct {
		personAnonymous
		grade int
	}{
		{personAnonymous: personAnonymous{"wick", 21}, grade: 2},
		{personAnonymous: personAnonymous{"ethan", 22}, grade: 3},
		{personAnonymous: personAnonymous{"bond", 21}, grade: 3},
	}

	for _, student := range sSliceAnonymousStruct {
		fmt.Println(student)
	}
	//
	/*
		A.24.11. Deklarasi Anonymous Struct Menggunakan Keyword var
		Cara lain untuk deklarasi anonymous struct adalah dengan menggunakan keyword var.
	*/
	fmt.Println("A.24.11. Deklarasi Anonymous Struct Menggunakan Keyword var\n")
	var student struct {
		person
		grade int
	}

	student.person = person{"wick", 21}
	student.grade = 2
	/*		Statement type student struct adalah contoh cara deklarasi struct. Maknanya akan berbeda ketika keyword type diganti var, seperti pada contoh di atas var student struct, yang artinya dicetak sebuah objek dari anonymous struct kemudian disimpan pada variabel bernama student.

			Deklarasi anonymous struct menggunakan metode ini juga bisa dilakukan beserta inisialisasi-nya.
	*/
	// hanya deklarasi
	var sAnonyVar1 struct {
		grade int
	}

	// deklarasi sekaligus inisialisasi
	var sAnonyVar2 = struct {
		grade int
	}{
		12,
	}
	sAnonyVar1.grade = 21
	sAnonyVar2.grade = 2
	fmt.Printf("sAnonyVar1: %v, sAnonyVar2: %v", sAnonyVar1, sAnonyVar2)
	/*
		A.24.15. Type Alias
		Sebuah tipe data, seperti struct, bisa dibuatkan alias baru, caranya dengan type NamaAlias = TargetStruct. Contoh:
	*/
	type PersonAlias struct {
		name string
		age  int
	}
	type People = PersonAlias

	var p1 = PersonAlias{"wick", 21}
	fmt.Println(p1)

	var p2 = PersonAlias{"wick", 21}
	fmt.Println(p2)
	/*
		Pada kode di atas, sebuah alias bernama People dibuat untuk struct Person.

		Casting dari objek (yang dicetak lewat struct tertentu) ke tipe yang merupakan alias dari struct pencetak, hasilnya selalu valid. Berlaku juga sebaliknya.
	*/
	people2 := People{"wick", 21}
	fmt.Println(PersonAlias(people2))

	person2 := PersonAlias{"wick", 21}
	fmt.Println(People(person2))

	type People1 struct {
		name string
		age  int
	}
	type People2 = struct {
		name string
		age  int
	}

	type Number = int
	var num Number = 12
	fmt.Println(num)
}

/*
A.24.12. Nested struct
Nested struct adalah anonymous struct yang di-embed ke sebuah struct.
Deklarasinya langsung di dalam struct peng-embed. Contoh:
*/
type structNested struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

/*
A.24.13. Deklarasi Dan Inisialisasi Struct Secara Horizontal
Deklarasi struct bisa dituliskan secara horizontal, caranya bisa dilihat pada kode berikut:

// di goland langsung berubah
type person struct { name string; age int; hobbies []string }
var p1 = struct { name string; age int } { age: 22, name: "wick" }
var p2 = struct { name string; age int } { "ethan", 23 }
*/
/*
A.24.14. Tag property dalam struct
Tag merupakan informasi opsional yang bisa ditambahkan pada masing-masing property struct.
*/
type personTagProperty struct {
	name string `tag1`
	age  int    `tag2`
}

//Tag biasa dimanfaatkan untuk keperluan encode/decode data json. Informasi tag juga bisa diakses lewat reflect. Nantinya akan ada pembahasan yang lebih detail mengenai pemanfaatan tag dalam struct, terutama ketika sudah masuk chapter JSON.

// Anonymous Struct
type personAnonymous struct {
	name string
	age  int
}

// Embedded Struct
type person struct {
	name string
	age  int
}

type student2 struct {
	grade int
	person
}

/*
	Disiapkan struct person dengan properti yang tersedia adalah name dan age.
	Disiapkan struct student2 dengan property grade.
	Struct person di-embed ke dalam struct student2.
	Caranya cukup mudah, yaitu dengan menuliskan nama struct yang ingin di-embed ke dalam body struct target.
*/

// Embedded Struct Dengan Nama Property Yang Sama
type person3 struct {
	name string
	age  int // age of person
}

type student3 struct {
	person3
	age   int // age of student
	grade int
}
