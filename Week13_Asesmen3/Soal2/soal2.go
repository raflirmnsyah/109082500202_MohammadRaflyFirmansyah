package main

import "fmt"

const NMAX = 100

type Pemain struct {
	nama   string
	gol    int
	assist int
}

type ArrPemain [NMAX]Pemain

func SelectionSort(A *ArrPemain, n int) {
	var idx int
	for i := 0; i < n-1; i++ {
		idx = i
		for j := i + 1; j < n; j++ {
			if A[j].gol > A[idx].gol ||
				(A[j].gol == A[idx].gol && A[j].assist > A[idx].assist) {
				idx = j
			}
		}
		A[i], A[idx] = A[idx], A[i]
	}
}

func main() {
	var A ArrPemain
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&A[i].nama, &A[i].gol, &A[i].assist)
	}

	SelectionSort(&A, n)

	fmt.Println()
	fmt.Println("Hasil Sorting")

	for i := 0; i < n; i++ {
		fmt.Println(A[i].nama, A[i].gol, A[i].assist)
	}
}