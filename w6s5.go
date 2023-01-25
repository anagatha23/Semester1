package main

import "fmt"

func main(){
	var angka, digit1, digit2, selisih1, selisih2 int
	var apakah1 bool

	fmt.Scan(&angka)
	
	

	for angka > 0{
		digit1 = angka%10
		angka = angka/10
		digit2 = angka%10
		angka = angka/10
		selisih1 = digit1 - digit2
		selisih2 = digit2 - digit1
		apakah1 = selisih1 == 1 || selisih2 == 1
		if selisih1 != 1 || selisih2 != 1{
			break
		}
	}
	
	fmt.Println(apakah1)
}

//Apabila didefinisikan sebuah bilangan konsekutif adalah bilangan yang selisih setiap digit yang bersebelahannya adalah satu.
//Maka buatlah algoritma untuk menentukan suatu bilangan konsekutif atau tidak.

//Masukan sebuah bilangan bulat positif X.

//Keluaran adalah sebuah nilai Boolean yang menyatakan X adalah konsekutif atau tidak.

//contoh masukkan: 1234321
//keluaran: true
//selisih setiap digit adalah 1