package main

import "fmt"

type arrIkan [1000]float64

func hitungWadah(ikan arrIkan, x, y int, total *arrIkan, jumlahWadah *int) {
	index := 0
	*jumlahWadah = x / y
	if x%y != 0 {
		*jumlahWadah++
	}
	for i := 0; i < *jumlahWadah; i++ {
		sum := 0.0
		for j := 0; j < y && index < x; j++ {
			sum += ikan[index]
			index++
		}
		total[i] = sum
	}
}

func rerata(total arrIkan, n int) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += total[i]
	}
	return sum / float64(n)
}

func main() {
	var x, y int
	var ikan, total arrIkan
	var jumlahWadah int
	fmt.Scan(&x, &y)
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}
	hitungWadah(ikan, x, y, &total, &jumlahWadah)
	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", total[i])
	}
	fmt.Println()
	fmt.Printf("%.2f\n", rerata(total, jumlahWadah))
}