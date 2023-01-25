//kelipatan
package main 
import "fmt"
func main(){
	var bil int
	fmt.Scan(&bil)
	if bil%3 == 0{
		fmt.Println("Kelipatan 3")
	} 
	if bil%5 == 0{
		fmt.Println("Kelipatan 5")
	}
	if bil%3 != 0 && bil%5 != 0{
		fmt.Println()
	}
}

//Buatlah program yang akan menentukan apakah sebuah bilangan adalah
//kelipatan 3 atau kelipatan 5 atau  kelipatan 3 dan juga kelipatan 5.

//Masukan adalah sebuah bilangan bulat positif.

//Keluaran berupa string yang menyatakan jenis bilangan tersebut, "Kelipatan 5", "Kelipatan 3" dan
//jika bukan kelipatan 3 atau 5, maka tidak menampilkan apapun.