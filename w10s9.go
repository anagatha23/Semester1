//akhir tahun
package main
import "fmt"
func main(){
	var belanja int
	var bersedia bool
	fmt.Scan(&belanja, &bersedia)
	fmt.Println("Kartu?", bersedia)
	if belanja >= 100000 && bersedia == true{
		fmt.Println("Diskon?", belanja >= 100000)
		fmt.Println("Cashback?", (belanja >= 200000) && (bersedia == true))
		fmt.Println("Rp.", (belanja*9/10) - 75000)
	}else if belanja >= 100000 && bersedia == false{
		fmt.Println("Diskon?", belanja >= 100000)
		fmt.Println("Cashback?", (belanja >= 200000) && (bersedia == true))
		fmt.Println("Rp.", (belanja*9/10))
	}else if belanja >= 100000 && belanja <= 200000 && bersedia == true{
		fmt.Println("Diskon?", belanja >= 100000)
		fmt.Println("Cashback?", (belanja >= 200000) && (bersedia == true))
		fmt.Println("Rp.", (belanja*9/10))
	}else if belanja <= 100000{
		fmt.Println("Diskon?", belanja >= 100000)
		fmt.Println("Cashback?", (belanja >= 200000) && (bersedia == true))
		fmt.Println("Rp.", belanja)
	}
}

//Sebuah mini market sedang merayakan acara akhir tahun, sehingga diberikanlah promo pada hari
//tersebut berupa cashback, diskon, dan juga kartu membership.
//Memperoleh kartu apabila pembeli bersedia dibuatkan dan memperoleh diskon.
//Diskon 10 % diberikan apabila belanja minimal Rp. 100.000, dan
//Cashback Rp. 75.000 diberikan apabila belanja minimal Rp. 200.000 dan memperoleh kartu.

//Masukan berupa bilangan bulat yang menyatakan total belanja dan Boolean yang menyatakan bersedia atau tidaknya untuk dibuatkan kartu.

//Keluaran terdiri dari empat baris, yang masing-masing barisnya merupakan nilai boolean yang menyatakan
//pembeli memperoleh kartu, diskon dan juga cashback. Baris terakhir adalah nominal belaja yang harus dibayar pembeli.

//contoh masukan
//1230000 false

//keluaran
//Kartu? false
//Diskon? true
//Cashback? false
//Rp. 1107000