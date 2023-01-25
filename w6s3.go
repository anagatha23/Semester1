package main

import "fmt"

func main(){
	var angka, digit, jumlah int

	fmt.Scan(&angka)
	jumlah = 0
	

	for angka > 0{
		digit = angka%10
		
		jumlah += digit
		angka = angka/10
		fmt.Print(digit, " ")
		
	}
	
	fmt.Println(jumlah)
}

//Sebuah algoritma digunakan untuk mencacah digit suatu bilangan bulat positif.

//Masukan terdiri dari sebuah bilangan bulat positif X.
//contoh 1234

//Keluaran teridri dari dua baris. Baris pertama adalah nilai setiap digit dari bilangan X dan dipisahkan oleh spasi.
//Baris kedua adalah total penjumlahan setiap digit dari X.
//4 3 2 1
//10