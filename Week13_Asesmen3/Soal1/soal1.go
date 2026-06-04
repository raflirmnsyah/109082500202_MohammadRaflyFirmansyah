package main

import "fmt"

const NMAX = 1000000

type arrInt [NMAX]int

func SelectionSort(T *arrInt, n int) {
	var pass, idx, i int

	for pass = 0; pass < n-1; pass++ {
		idx = pass
		for i = pass + 1; i < n; i++ {
			if T[i] < T[idx] {
				idx = i
			}
		}
		T[pass], T[idx] = T[idx], T[pass]
	}
}

func median(T arrInt, n int) float64 {
	if n%2 == 1 {
		return float64(T[n/2])
	}
	return float64((T[n/2-1] + T[n/2]) / 2)
}

func main() {
	var A arrInt
	var x int
	var n int = 0

	fmt.Scan(&x)

	for x != -5313541 && n < NMAX {
		if x == 0 {
			SelectionSort(&A, n)
			fmt.Println(int(median(A, n)))
		} else {
			A[n] = x
			n++
		}
		fmt.Scan(&x)
	}
}