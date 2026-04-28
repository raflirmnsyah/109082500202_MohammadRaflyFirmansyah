package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&N)

	var arr [100]int

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nSemua isi array:")
	for i := 0; i < N; i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\n\nElemen indeks ganjil:")
	for i := 1; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("\n\nElemen indeks genap:")
	for i := 0; i < N; i += 2 {
		fmt.Print(arr[i], " ")
	}

	var x int
	fmt.Print("\n\nMasukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Println("Elemen indeks kelipatan", x, ":")
	for i := 0; i < N; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	var idx int
	fmt.Print("\n\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&idx)

	for i := idx; i < N-1; i++ {
		arr[i] = arr[i+1]
	}
	N-- 

	fmt.Println("Array setelah dihapus:")
	for i := 0; i < N; i++ {
		fmt.Print(arr[i], " ")
	}

	var jumlah int
	for i := 0; i < N; i++ {
		jumlah += arr[i]
	}
	rata := float64(jumlah) / float64(N)
	fmt.Println("\n\nRata-rata:", rata)

	var total float64
	for i := 0; i < N; i++ {
		total += math.Pow(float64(arr[i])-rata, 2)
	}
	stdDev := math.Sqrt(total / float64(N))
	fmt.Println("Standar deviasi:", stdDev)

	var cari, frek int
	fmt.Print("\nMasukkan angka yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)

	for i := 0; i < N; i++ {
		if arr[i] == cari {
			frek++
		}
	}

	fmt.Println("Frekuensi:", frek)
}