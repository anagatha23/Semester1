//konsonan
package main
import "fmt"
func main(){
	var k int
	fmt.Scanf("%c", &k)
	if k == 'A' || k == 'I' || k == 'U' || k == 'E' || k == 'O' || k == 'a' || k == 'i' || k == 'u' || k == 'e' || k == 'o'{
		fmt.Print("bukan konsonan")
	}else{
		fmt.Print("konsonan")
	}
}

//Sebuah program digunakan untuk menentukan karakter yang diberikan adalah huruf konsonan atau bukan.

//Masukan terdiri dari satu karakter huruf.

//Keluaran terdiri dari teks "konsonan" apabila karakter adalah huruf konsonan
//atau "bukan konsonan" untuk kemungkinan huruf yang lain.