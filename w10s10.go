//avatar
package main
import "fmt"
func main(){
	var penumpang, dewasa, kecil, takBerangkat int
	fmt.Scan(&penumpang)
	if penumpang <= 15 && penumpang%5 != 0{
		dewasa =  penumpang/5 + 1
		fmt.Print("Dewasa ", dewasa)
	}else if penumpang <= 15 && penumpang%5 == 0{
		dewasa =  penumpang/5
		fmt.Print("Dewasa ", dewasa)
	}else if penumpang > 15 && penumpang <= 25{
		dewasa = 3
		if (penumpang-15)%2 != 0{ 
			kecil = (penumpang-15)/2 + 1
		}else{
			kecil = (penumpang-15)/2
		}
		fmt.Print("Dewasa ", dewasa,", kecil ", kecil)
	}else{
		dewasa = 3
		kecil = 5
		takBerangkat = penumpang%25
		fmt.Print("Dewasa ", dewasa,", kecil ", kecil, ", dan ", takBerangkat, " tak berangkat")
	}
}

//Appa adalah sejenis sky bison yang digunakan untuk transportasi dalam serial  Avatar: The Last Airbender.
//Pada kisah kali ini Avatar Aang sedang bersiap untuk menyerang negara Api, untuk itu dia harus menyediakan
//beberapa Appa untuk mengangkut rekan-rekannya. Aang memiliki tiga ekor Appa dewasa dan lima Appa kecil.
//Appa dewasa dapat mengangkut lima orang, sedangkan Appa kecil hanya 2 orang perekornya.

//Masukan adalah sebuah bilangan bulat asli p yang menyatakan jumlah orang yang akan berangkat.

//Keluaran berupa banyak Appa dewasa dan Appa kecil yang digunakan. Tampilkan juga jumlah orang
//yang tidak bisa berangkat karena telah kehabisan Appa apabila ada.

//Catatan : Appa dewasa lebih diprioritaskan dibandingkan yang kecil, dan jumlah dan kapasitas setiap jenis Appa adalah konstanta.