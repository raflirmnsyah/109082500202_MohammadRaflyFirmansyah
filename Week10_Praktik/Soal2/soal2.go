package main

import "fmt"

const NMAX int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrMahasiswa [NMAX]mahasiswa

func isiData(A *arrMahasiswa, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Print("Masukkan data ke-", i+1, "\n")
		fmt.Print("NIM : ")
		fmt.Scan(&A[i].NIM)
		fmt.Print("Nama : ")
		fmt.Scan(&A[i].nama)
		fmt.Print("Nilai : ")
		fmt.Scan(&A[i].nilai)
	}
}

func rataNilai(A arrMahasiswa, n int, nim string) float64 {
	var i, jumlah, count int
	for i = 0; i < n; i++ {
		if A[i].NIM == nim {
			jumlah += A[i].nilai
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return float64(jumlah) / float64(count)
}

func nilaiTerbesar(A arrMahasiswa, n int, nim string) int {
	var i, max int
	var found bool = false
	for i = 0; i < n; i++ {
		if A[i].NIM == nim {
			if !found || A[i].nilai > max {
				max = A[i].nilai
				found = true
			}
		}
	}
	return max
}

func main() {
	var data arrMahasiswa
	var n int
	var nimCari string
	fmt.Print("Jumlah data: ")
	fmt.Scan(&n)
	isiData(&data, n)
	fmt.Print("\nMasukkan NIM yang ingin dicari nilainya: ")
	fmt.Scan(&nimCari)
	fmt.Println("Nilai pertama dari NIM", nimCari, "adalah", rataNilai(data, n, nimCari))
	fmt.Println("Nilai terbesar dari NIM", nimCari, "adalah", nilaiTerbesar(data, n, nimCari))
}