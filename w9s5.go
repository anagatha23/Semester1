package main
import "fmt"
func main(){
	var jariA, jariB, jarak float64
	var tidakBeriris bool
	fmt.Scan(&jariA, &jariB, &jarak)
	tidakBeriris = jariA/2 + jariB/2 <= jarak
	fmt.Println(tidakBeriris)
}

//Buatlah program yang digunakan untuk menentukan apakah dua buah lingkaran menghasilkan irisan atau tidak.

//Masukan terdiri dari tiga nilai, yaitu jari-jari lingkaran A, B dan jarak titik pusat lingkarang A dan B

//Keluaran berupa boolean yang menyatakan lingkaran A dan B tidak beririsan.