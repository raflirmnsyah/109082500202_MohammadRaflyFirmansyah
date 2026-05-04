package main

import "fmt"

type arrBerat [1000]float64

func hitungMinMax(data arrBerat, n int, min, max *float64) {
	*min = data[0]
	*max = data[0]

	for i := 1; i < n; i++ {
		if data[i] < *min {
			*min = data[i]
		}
		if data[i] > *max {
			*max = data[i]
		}
	}
}

func main() {
	var n int
	var berat arrBerat
	var min, max float64

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&berat[i])
	}

	hitungMinMax(berat, n, &min, &max)

	fmt.Printf("%.2f %.2f\n", min, max)
}