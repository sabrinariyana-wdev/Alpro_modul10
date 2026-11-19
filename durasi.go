package main
import "fmt"

func main() {
    var a, b, durasi int

    fmt.Scan(&a, &b)
    if b >= a {
        durasi = b - a
    } else {
        durasi = (24 - a) + b
    }
    fmt.Println(durasi)
}

// program Durasi
// kamus
//     a, b, durasi : integer  
// algoritma
//     input (a, b)
//     if b >= a then  
//     durasi = b - a
//     else
//     durasi = (24 - a) + b
//     endif
//     output (durasi)
// endprogram