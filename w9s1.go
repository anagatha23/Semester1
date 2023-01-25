package main
import "fmt"
func main(){
	var tahun int
	fmt.Scan(&tahun)
	fmt.Println(((tahun%400 == 0) || (tahun%4 == 0)) && (tahun%100 != 0))
}

//Tahun disebut kabisat bukan hanya habis dibagi 4 karena tahun yang habis dibagi 100 belum tentu kabisat.
//Tahun kabisat adalah tahun yang habis dibagi 400 atau habis dibagi 4 tetapi tidak habis dibagi 100. 
//Contoh tahun kabisat adalah 2016 dan 2000. Tahun 2018 dan 2200 bukanlah tahun kabisat.

//Masukan berupa sebuah bilangan yang menyatakan tahun.

//Keluaran berupa sebuah boolean yang menyatakan tahun adalah kabisat.