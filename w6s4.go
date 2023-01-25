package main
import "fmt"

func main(){
	var n, m, x, y, bagiGula, bagiKopi int
	fmt.Scan(&n, &m, &x, &y)
	bagiGula = n/x
	bagiKopi = m/y

	if bagiGula <= bagiKopi {
		fmt.Println(bagiGula)
	} else {
		fmt.Println(bagiKopi)
	}

}

//Buatlah sebuah program untuk menghitung banyak cangkir kopi yang bisa dibuat apabila terdefinisi sejumlah n gula dan m kopi.
//Satu cangkir kopi memerlukan sejumlah x gula dan y kopi.

//Masukan berupa empat bilangan bulat yang dipisahkan spasi, n, m, x dan y. Di mana nilai x <= n dan y <= m.

//Keluaran berupa bilangan bulat yang menyatakan banyaknya cangkir kopi yang berhasil dibuat.

//contoh masukan: 5 9 1 3
//keluaran: 3
//Tersedia 5 gula dan 9 kopi, 1 cangkir memerlukan 1 gula dan 3 kopi, sehingga didapat 3 cangkir.