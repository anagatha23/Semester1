//parkir
package main
import "fmt"
func main(){
	var h1, m1, h2, m2, jam, menit int
	fmt.Scan(&h1, &m1, &h2, &m2)
	if h1 > h2 && m1 > m2{
		jam = 12 - h1 + h2 - 1
		menit = 60 - m1 + m2
	}else if h1 > h2{
		jam = 12 - h1 + h2
		menit = m2 - m1
	}else{
		jam = h2 - h1
		menit = m2 - m1
	}

	fmt.Println(jam, "jam", menit, "menit")
}

//Sebuah tempat wisata buka jam 7:00 pagi hingga 5:00 sore. Buatlah program yang digunakan untuk
//menghitung durasi parkir suatu kendaraan pada tempat wisata tersebut.

//Masukan terdiri dari empat nilai h1, m1 yang menyatakan jam dan menit kendaraan mulai parkir, dan h2, m2
//yang menyatakan jam dan menit kendaraan selesai parkir.

//Keluaran terdiri dari dua bilangan hh dan mm yang menyatakan durasi parkir dalam jam dan menitnya.