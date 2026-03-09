package main
import "fmt"

func main() {
	var berat int
	var biayaGram int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&berat)
	
	kg := berat / 1000
	gram := berat % 1000

	biayaKg := kg * 10000

	if gram >= 500 {
		biayaGram = 2500
	} else {
		biayaGram = gram * 15
	}

	total := biayaKg + biayaGram

	fmt.Println("\n===== Detail Perhitungan =====")
	fmt.Println("Detail berat :", kg, "kg +", gram, "gram")
	fmt.Println("Detail biaya : Rp.", biayaKg, "+ Rp.", biayaGram)
	fmt.Println("Total biaya: Rp", total)
}