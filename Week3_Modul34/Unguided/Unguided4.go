package main
import (
	"fmt"
	"math"
)

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi

	fmt.Println("Luas Persegi:", luas)
	fmt.Println("Keliling Persegi:", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println("Luas Persegi Panjang:", luas)
	fmt.Println("Keliling Persegi Panjang:", keliling)
}

func hitungLingkaran(r float64) {
	luas := math.Pi * r * r
	keliling := 2 * math.Pi * r

	fmt.Println("Luas Lingkaran:", luas)
	fmt.Println("Keliling Lingkaran:", keliling)
}

func main() {
	var pilihan int

	fmt.Println("1. Persegi")
	fmt.Println("2. Persegi Panjang")
	fmt.Println("3. Lingkaran")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&pilihan)

	switch pilihan {

	case 1:
		var sisi int
		fmt.Print("Masukkan sisi: ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)

	case 2:
		var panjang, lebar int
		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)

	case 3:
		var r float64
		fmt.Print("Masukkan jari-jari: ")
		fmt.Scan(&r)
		hitungLingkaran(r)

	default:
		fmt.Println("Menu tidak tersedia")
	}
}