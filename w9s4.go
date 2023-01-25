package main
import "fmt"
func main(){
	var x, y, z int
	var midPoint bool
	fmt.Scan(&x, &y, &z)
	midPoint = ((x+y)/2 == z) || ((x+z)/2 == y) || ((z+y)/2 == x)
	fmt.Println(midPoint)
}

//Buatlah program yang digunakan untuk menentukan apakah suatu bilangan
//adalah nilai tengah dari dua bilangan tertentu.

//Masukan terdiri dari tiga bilangan bulat.

//Keluaran berupa boolean yang menyatakan salah satu bilangan adalah nilai tengah dua bilangan lainnya.