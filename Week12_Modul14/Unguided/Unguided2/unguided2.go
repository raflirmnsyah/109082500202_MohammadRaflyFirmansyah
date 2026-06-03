package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&m)

		var ganjil [100000]int
		var genap [100000]int
		jmlGanjil := 0
		jmlGenap := 0

		for j := 0; j < m; j++ {
			var x int
			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[jmlGanjil] = x
				jmlGanjil++
			} else {
				genap[jmlGenap] = x
				jmlGenap++
			}
		}

		for j := 0; j < jmlGanjil-1; j++ {
			min := j
			for k := j + 1; k < jmlGanjil; k++ {
				if ganjil[k] < ganjil[min] {
					min = k
				}
			}
			temp := ganjil[j]
			ganjil[j] = ganjil[min]
			ganjil[min] = temp
		}

		for j := 0; j < jmlGenap-1; j++ {
			max := j
			for k := j + 1; k < jmlGenap; k++ {
				if genap[k] > genap[max] {
					max = k
				}
			}
			temp := genap[j]
			genap[j] = genap[max]
			genap[max] = temp
		}

		pertama := true

		for j := 0; j < jmlGanjil; j++ {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[j])
			pertama = false
		}

		for j := 0; j < jmlGenap; j++ {
			if !pertama {
				fmt.Print(" ")
			}
			fmt.Print(genap[j])
			pertama = false
		}

		fmt.Println()
	}
}