package main
import "fmt"

func main(){
	var tangki, ember, jumlah int
	var penuh bool

	fmt.Scan(&tangki)
	jumlah = 0

	for jumlah <= tangki{
		fmt.Scan(&ember)
		jumlah += ember
		penuh = jumlah >= tangki
		fmt.Println(penuh)
		if jumlah >= tangki {
			break
		}
	}
}

//Sebuah tanki kosong dengan kapasitas T liter akan diisi air menggunakan ember E dengan berbagai ukuran volume (0 < E ≤ N).
//Buatlah algoritma untuk mengisi tanki dari asalnya kosong hingga penuh. 

//Masukan terdiri dari beberapa baris. Baris pertama adalah sebuah bilangan bulat positif T yang menyatakan kapasitas tanki.
//T baris berikutnya, masing-masing adalah bilangan bulat V yang menyatakan volume air dalam ember E dimasukkan ke dalam tanki.

//Keluaran terdiri dari beberapa baris, yang masing-masing baris berisi Boolean yang menyatakan
//tanki penuh atau tidak setiap kali pengisian tanki.

//contoh masukan
//30
//5
//10
//5
//5
//10
//keluaran
//false
//false
//false
//false
//true

//T = 30
//5 + 10 + 5 + 5 + 10 = 35