package main
import "fmt"

func main() {
    var un, pw string
    var try int
    
    
    
    
    try = 0
    
    for {
		fmt.Scan(&un, &pw)
        if un == "admin" && pw == "admin" {
            break
        }
        try++
    }
    fmt.Println(try, "Login berhasil")
}

//Sebuah program digunakan untuk menghitung berapa banyak seorang user melakukan gagal login karena salah input username dan password.

//Masukan terdiri dari username dan password. Apabila username dan password salah, maka lakukan proses input kembali sampai username
//dan password benar. Asumsi username dan password yang benar adalah "admin" dan "admin"

//Keluaran terdiri dari berapa banyak user gagal melakukan login, dan sebuah pesan "Login berhasil".