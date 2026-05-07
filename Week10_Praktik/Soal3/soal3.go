package main

import "fmt"

const nProv int = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Print("Masukkan data ke-", i+1, " : ")
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	maxIdx := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	var hasil float64
	fmt.Println("\n--- Prediksi Jumlah Penduduk Tahun Depan ---")
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			hasil = (tumbuh[i] + 1) * float64(pop[i])
			fmt.Println(prov[i], int(hasil))
		}
	}
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func main() {
	var prov NamaProv
	var pop PopProv
	var tumbuh TumbuhProv
	var cari string
	var idx int
	fmt.Println("=== Masukkan Nama Provinsi, Populasi, Pertumbuhan ===")
	InputData(&prov, &pop, &tumbuh)
	fmt.Print("\nMasukkan nama provinsi yang dicari : ")
	fmt.Scan(&cari)
	idx = ProvinsiTercepat(tumbuh)
	fmt.Println("\nProvinsi dengan angka pertumbuhan tercepat :", prov[idx])
	fmt.Println("Indeks provinsi :", IndeksProvinsi(prov, cari))
	Prediksi(prov, pop, tumbuh)
}