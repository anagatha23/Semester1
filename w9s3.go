package main
import "fmt"
func main(){
	var belanja int
	var sedia bool
	fmt.Scan(&belanja, &sedia)
	fmt.Println("Kartu?", ((belanja > 100000) && (sedia)))
	fmt.Println("Diskon?", belanja > 100000)
	fmt.Println("Cashback?", belanja > 200000)

}

//Sebuah mini market sedang merayakan ulang tahunnya, sehingga diberikanlah promo pada hari tersebut
//berupa cashback, diskon, dan juga kartu membership.
//Memperoleh kartu apabila pembeli bersedia dibuatkan dan memperoleh diskon. 
//Diskon diberikan apabila belanja minimal Rp. 100.000.
//Cashback diberikan apabila belanja minimal Rp. 200.000 dan memperoleh kartu.

//Masukan berupa bilangan bulat yang menyatakan total belanja dan Boolean
//yang menyatakan bersedia atau tidaknya untuk dibuatkan kartu.

//Keluaran berupa tiga baris, yang masing-masing barisnya merupakan nilai Boolean
//yang menyatakan pembeli memperoleh kartu, diskon dan juga cashback.

//Contoh masukan: 1230000 false

//keluaran:
//Kartu? false
//Diskon? true
//Cashback? false