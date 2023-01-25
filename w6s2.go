package main
import "fmt"

func main(){
	var saldo, trans int

	saldo = 0
	fmt.Scan(&trans)

	for trans != 0{
		
		saldo+= trans
		fmt.Scan(&trans)


	}
	fmt.Println(saldo)
}

//Buatlah algoritma untuk menghitung saldo uang yang ada dalam dompet seorang mahasiswa di akhir bulan.

//Masukan terdiri dari serangkaian bilangan bulat yang menyatakan transaksi keluar masuk dompet.
//Bilangan tanpa tanda berarti jumlah uang yang masuk ke dalam dompet, bilangan negatif menyatakan uang keluar dari dompet.
//Masukan berakhir apabila jumlah uang yang diberikan adalah 0 (nol)

//Keluaran adalah jumlah saldo uang dalam dompet.