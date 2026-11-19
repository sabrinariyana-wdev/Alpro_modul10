package main
import "fmt"

func main() {
    var j, tarif int

	fmt.Scan(&j)
	tarif = 10000
	if j > 5{
		tarif += 4000
	}
	fmt.Println(tarif)
}

// program Tarif
// kamus
// 	j, tarif : integer
// algoritma
// 	input (j)
// 	tarif = 10000
// 	if j > 5 then
// 		tarif = tarif + 4000
// 	endif
// 	output (tarif)
// endprogram