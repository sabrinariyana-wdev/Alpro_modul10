package main
import "fmt"

func main() {
    var s int

    fmt.Scan(&s)
    if s >= 450 {
        fmt.Println("KELAS REGULER")
    } else {
        fmt.Println("KELAS MATRIKULASI")
    }
}

// program Kelas
// kamus
//     s : integer
// algoritma
//     input (s)
//     if s >= 450 then
//         output ("KELAS REGULER")    
//     else
//         output ("KELAS MATRIKULASI")
//     endif
// endprogram