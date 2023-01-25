//profit
package main
import "fmt"
func main(){
	var untungSebelum, untungSekarang float64
	fmt.Scan(&untungSebelum, &untungSekarang)
	if untungSebelum > untungSekarang{
		fmt.Println("Penurunan sebesar", untungSebelum - untungSekarang)
	}else if untungSebelum < untungSekarang{
		fmt.Println("Peningkatan sebesar", untungSekarang - untungSebelum)
	}else{
		fmt.Println("Tetap")
	}
}

//Seorang pedagang sedang menghitung keuntungan yang diperoleh bulan ini dibandingkan bulan sebelumnya.

//Masukan terdiri dari bilangan riil yang menyatakan keuntungan bulan ini dan bulan sebelumnya.

//Keluaran terdiri dari string yang menyatakan peningkatan atau penurunan keuntungan sebesar x.