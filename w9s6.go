package main
import "fmt"

func main(){
	var i, n int
	var valid bool
	var item1, item2, item3, item4, item5 bool
	fmt.Scan(&n)
	
	for i = 1; i <= n;{
		fmt.Scan(&item1, &item2, &item3, &item4, &item5)
		valid = (item1 == true) && (item2 == true) && (item3 == true) && (item4 == true) && (item5 == true)
		i = i + 1
		if valid != true{
			break
		}
	}
	fmt.Println(valid)
}

//Anda adalah seorang ketua tim dalam kegiatan pramuka dan ditugaskan untuk mengecek kesiapan tim berupa kelengkapan anggotanya
//berjumlah 𝑁 orang. Terdapat 5 item yang harus dicek kelengkapannya pada setiap anggota
//Setelah selesai laporkan kepada pembina pramuka terkait kesiapan tim anda.

//Masukan terdiri dari beberapa baris. Baris pertama adalah sebuah bilangan bulat positif N
//yang menyatakan jumlah anggota tim. 𝑁 baris
//berikutnya masing-masing berisistatus boolean dari kelengkapan 5 item setiap anggota.

//Keluaran adalah sebuahboolean yangmenyatakan status kesiapan tim.

