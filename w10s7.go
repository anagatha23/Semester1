//liga sepakbola
package main
import "fmt"
func main(){
	var g1, g2, g3, g4 int
	fmt.Scan(&g1, &g2, &g3, &g4)
	if g1 >= g2 && g1 >= g3 && g1 >= g4{
		fmt.Print(g1)
	}else if g2 >= g1 && g2 >= g3 && g2 >= g4{
		fmt.Print(g2)
	}else if g3 >= g1 && g3 >= g2 && g3 >= g4{
		fmt.Print(g3)
	}else if g4 >= g1 && g4 >= g2 && g4 >= g3{
		fmt.Print(g4)
	}
	if g1 <= g2 && g1 <= g3 && g1 <= g4{
		fmt.Print(" ", g1)
	}else if g2 <= g1 && g2 <= g3 && g2 <= g4{
		fmt.Print(" ", g2)
	}else if g3 <= g1 && g3 <= g2 && g3 <= g4{
		fmt.Print(" ", g3)
	}else if g4 <= g1 && g4 <= g2 && g4 <= g3{
		fmt.Print(" ", g4)
	}
}

//Buatlah program yang digunakan untuk menentukan jumlah gol terbanyak dan jumlah gol tersedikit dari suatu grup liga sepak bola.

//Masukan terdiri dari empat bilangan yang menyatakan jumlah gol yang berhasil dicetak empat tim bola dalam suatu grup.

//Keluaran terdiri dari dua nilai yang menyatakan jumlah gol terbanyak dan jumlah gol tersedikit.