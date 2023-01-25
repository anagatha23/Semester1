//Temperatur
package main
import "fmt"
func main(){
	var t1, t2, t3, t4, t5 float64
	fmt.Scan(&t1, &t2, &t3, &t4, &t5)
	if t1 < t2 && t2 < t3 && t3 < t4 && t4 < t5{
		fmt.Println("Stabil naik")
	}else if t1 > t2 && t2 > t3 && t3 > t4 && t4 > t5{
		fmt.Println("Stabil turun")
	}else{
		fmt.Println("Tidak stabil")
	}
}

//Sebuah sensor digunakan untuk mencatat perubahan temperatur suatu zat radioaktif X. 
//Buatlah program yang digunakan untuk pencatatan temperature pada zat radioaktif X tersebut.

//Masukan terdiri dari lima bilangan riil yang menyatakan catatan temperatur zat radioaktif X dalam 5 kali pencatatan.

//Keluaran terdiri dari string "Stabil naik", atau "Stabil turun", atau
//"Tidak stabil" berdasarkan perubahan temperature pada lima kali pencatatan.

