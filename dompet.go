package main
import "fmt"

func main() {
	var s,p int

	fmt.Scan(&s,&p)
	if (s >= 0 && s <= 100000) || (p >= 0 && p <= 100000) {
        if s > p {
            fmt.Println("Bisa beli")
        } else if s == p {
            fmt.Println("Bisa pas")
        } else {
            fmt.Println("Tidak bisa beli")
        }
    }
}

// program Dompet
// kamus
// 	s, p : integer
// algoritma
// 	input (s,p)
// 	if ( s >= 0 and s <= 100000) or (p >= 0 and p <= 100000) then
// 		if s > p then
// 			output ("Bisa beli")
// 		else if s==p then
// 			output ("Bisa pas")
// 		else
// 			output ("Tidak bisa beli")
// 		endif
// 	endif
// endprogram