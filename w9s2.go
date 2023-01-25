package main
import "fmt"
func main(){
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println((a + b > c) && (b + c > a) && (c + a > b))
}

//Buatlah program yang digunakan untuk menentukan apakah nilai tiga bilangan dapat membentuk suatu segitiga. 

//Masukan terdiri dari tiga bilangan bulat.

//Keluaran adalah boolean yang menyatakan tiga bilangan bulat dapat membentuk segitiga