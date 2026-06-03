package main

import "fmt"

const NMAX = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [NMAX]Buku

func BacaBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)

	for i := 0; i < *n; i++ {
		fmt.Scan(
			&pustaka[i].id,
			&pustaka[i].judul,
			&pustaka[i].penulis,
			&pustaka[i].penerbit,
			&pustaka[i].eksemplar,
			&pustaka[i].tahun,
			&pustaka[i].rating,
		)
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	var j int

	for i := 1; i < n; i++ {
		temp = pustaka[i]
		j = i - 1

		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}

		pustaka[j+1] = temp
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		fmt.Println(
			pustaka[0].judul,
			pustaka[0].penulis,
			pustaka[0].penerbit,
			pustaka[0].tahun,
		)
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5
	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Println(
			pustaka[i].judul,
			pustaka[i].penulis,
			pustaka[i].penerbit,
			pustaka[i].tahun,
		)
	}
}

func CariBuku(pustaka DaftarBuku, n int, rating int) {
	kiri := 0
	kanan := n - 1
	ketemu := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if pustaka[tengah].rating == rating {
			ketemu = tengah
			break
		} else if rating > pustaka[tengah].rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu == -1 {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	} else {
		fmt.Println(
			pustaka[ketemu].judul,
			pustaka[ketemu].penulis,
			pustaka[ketemu].penerbit,
			pustaka[ketemu].tahun,
			pustaka[ketemu].rating,
		)
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	var ratingCari int

	BacaBuku(&pustaka, &n)

	UrutBuku(&pustaka, n)

	CetakTerfavorit(pustaka, n)

	Cetak5Terbaru(pustaka, n)

	fmt.Scan(&ratingCari)

	CariBuku(pustaka, n, ratingCari)
}