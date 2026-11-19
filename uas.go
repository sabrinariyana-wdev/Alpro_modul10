package main
import "fmt"

func main() {
	var t,h,k int
	var hasil string

	fmt.Scan(&t,&h,&k)
	hasil = "tidak boleh UAS"
	if h*100 >= 75*t && k >= 8 {
		hasil = "boleh UAS"
	}
	fmt.Println(hasil)
}

// program UASQ
// kamus
// 	t,h,k : integer
// 	hasil : string
// algoritma
// 	input (t,h,k)
// 	hasil = "tidak boleh UAS"
// 	if h*100 >= 75*t and k >= 8 then
// 		hasil = "boleh UAS"
// 	endif
// 	output (hasil)
// endprogram