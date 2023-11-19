package main

import (
	"fmt"
	"strings"
)

/*
A.25. Method
Method adalah fungsi yang menempel pada type (bisa struct atau tipe data lainnya). Method bisa diakses lewat variabel objek.

Keunggulan method dibanding fungsi biasa adalah memiliki akses ke property struct hingga level private (level akses nantinya akan dibahas lebih detail pada chapter selanjutnya). Dan juga, dengan menggunakan method sebuah proses bisa di-enkapsulasi dengan baik.
*/
type sMethod struct {
	name  string
	grade int
}

func (s sMethod) sayHello() {
	fmt.Println("halo", s.name)
}
func (s sMethod) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
} /*
func (s student) sayHello() maksudnya adalah fungsi sayHello dideklarasikan sebagai method milik struct student. Pada contoh di atas struct student memiliki dua buah method, yaitu sayHello() dan getNameAt().
*/

func (s sMethod) chName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *sMethod) chName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {
	var s1 = sMethod{"muh syahrur", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan :", name)

	//Method Pointer
	var s2 = sMethod{"muh syahrur", 21}
	fmt.Println("s1 before", s1.name)
	// john wick

	s2.chName1("john all")
	fmt.Println("s1 after chName1", s1.name)

	s2.chName2("ethan hunt")
	fmt.Println("s1 after chName2", s1.name)
	/*Setelah eksekusi statement s2.changeName1("jason bourne"), nilai s2.name tidak berubah. Sebenarnya nilainya berubah tapi hanya dalam method changeName1() saja, nilai pada reference di objek-nya tidak berubah. Karena itulah ketika objek di print value dari s1.name tidak berubah.

	  Keistimewaan lain method pointer adalah, method itu sendiri bisa dipanggil dari objek pointer maupun objek biasa.*/
	// pengaksesan method dari variabel objek biasa
	var s3 = sMethod{"john wick", 21}
	s3.sayHello()

	// pengaksesan method dari variabel objek pointer
	var s4 = &sMethod{"ethan hunt", 22}
	s4.sayHello()
	/*
		Penggunaan Fungsi strings.Split()
		Pada chapter ini ada fungsi baru yang kita gunakan: strings.Split(). Fungsi ini berguna untuk memisah string menggunakan pemisah yang ditentukan sendiri. Hasilnya adalah array berisikan kumpulan substring.

			strings.Split("ethan hunt", " ")
		// ["ethan", "hunt"]

	*/

}
