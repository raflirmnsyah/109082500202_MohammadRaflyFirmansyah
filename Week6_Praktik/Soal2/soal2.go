package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
	return 0
}

func biayaPerHari(jumlahMhs int) int {
	biayaMakan := 2 * 35000       
	biayaPenginapan := 250000
	uangSaku := 300000
	totalPerOrang := biayaMakan + biayaPenginapan + uangSaku
	return jumlahMhs * totalPerOrang
}

func perhitunganBiaya(jumlahMhs int, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hari := tanggunganHari(lamaPerjalanan, tujuan)
	biayaHarian := biayaPerHari(jumlahMhs)
	total := float64(hari * biayaHarian)
	if tujuan == "mancanegara" {
		total = total * 1.5
	}
	*totalBiaya = total
}

func main() {
	var jumlahMhs int
	var lamaPerjalanan int
	var tujuan string
	var totalBiaya float64
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&jumlahMhs)
	fmt.Print("Lama study tour (hari): ")
	fmt.Scan(&lamaPerjalanan)
	fmt.Print("Masukkan tujuan (domestik/mancanegara): ")
	fmt.Scan(&tujuan)
	perhitunganBiaya(jumlahMhs, lamaPerjalanan, tujuan, &totalBiaya)
	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U = Rp. %.0f\n", totalBiaya)
}