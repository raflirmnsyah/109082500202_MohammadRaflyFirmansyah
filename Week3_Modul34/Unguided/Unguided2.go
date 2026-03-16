package main
import "fmt"
func hitungBiaya(jenis string, masuk, keluar int) int {
	jam := keluar - masuk
	biaya := 0

	if jenis == "motor" {
		for i := 0; i < jam; i++ {
			if masuk+i <= 17 {
				biaya += 4000
			} else {
				biaya += 5000
			}
		}
	} else if jenis == "mobil" {
		for i := 0; i < jam; i++ {
			if masuk+i <= 17 {
				biaya += 6000
			} else {
				biaya += 7000
			}
		}
	}

	return biaya
}

func main() {
	var jenis string
	var masuk, keluar int
	total := 0

	for {
		fmt.Print("Jenis kendaraan: ")
		fmt.Scan(&jenis)

		if jenis == "-" {
			break
		}

		fmt.Print("Jam masuk: ")
		fmt.Scan(&masuk)

		fmt.Print("Jam keluar: ")
		fmt.Scan(&keluar)

		biaya := hitungBiaya(jenis, masuk, keluar)
		fmt.Println("Biaya parkir:", biaya)

		total += biaya
	}

	fmt.Println("Total pendapatan parkir:", total)
}