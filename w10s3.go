//segitiga
package main
import "fmt"
func main(){
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a == b && a == c && c == b{
		fmt.Println("Segitiga Sama Sisi")
	}else if (a == b && a != c) || (a == c && a != b) || (b == c && b != a){
		fmt.Println("Segitigas Sama Kaki")
	}else{
		fmt.Println("Segitigas Sembarang")
	}
}

//Buatlah program yang digunakan untuk menentukan jenis segitiga 
//(segitiga sama sisi, segitiga sama kaki, atau segitiga sembarang) berdasarkan nilai tiga sisinya. 

//Masukan terdiri dari tiga bilangan bulat positif yang menyatakan sisi dari suatu segitiga

//Keluaran adalah suatu string yang menyatakan jenis segitiga:
//"Segitiga Sama Sisi", "Segitiga Sama Kaki", dan "Segitiga Sembarang".