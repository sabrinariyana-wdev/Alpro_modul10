package main
import "fmt"

func main() {
	var daging, masakan, tambahan string

	fmt.Scan(&daging)
	if daging == "ayam" {
		fmt.Println("Sop atau Katsu")
		fmt.Scan(&masakan)

		if masakan == "sop" {
			tambahan = "Beli sayur"
		} else if masakan == "katsu" {
			tambahan = "Beli tepung"
		}

	} else if daging == "sapi" {
		fmt.Println("Teriyaki atau Semur")
		fmt.Scan(&masakan)

		if masakan == "teriyaki" {
			tambahan = "Beli bombay"
		} else if masakan == "semur" {
			tambahan = "Beli kecap"
		}
	}
	fmt.Println("Daging:", daging, "Tambahan:", tambahan)
}

// program Masakan
// kamus
// 	daging, masakan, tambahan : string
// algoritma
// 	input (daging)
// 	if daging == "ayam" then
// 		output ("Sop atau Katsu")
// 		input (masakan)
// 		if masakan == "sop" then	
// 			tambahan = "Beli sayur"
// 		else if masakan == "katsu" then
// 			tambahan = "Beli tepung"	
// 	else if daging == "sapi" then
// 		output ("Teriyaki atau Semur")
// 		input (masakan)
// 		if masakan == "Teriyaki" then	
// 			tambahan = "Beli bombay"
// 		else if masakan == "Semur" then
// 			tambahan = "Beli kecap"	
// 		endif
// 	endif
// endprogram