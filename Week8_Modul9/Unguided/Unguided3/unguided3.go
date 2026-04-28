package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var pemenang [100]string
	var skorA, skorB int
	i := 0

	for {
		fmt.Printf("Pertandingan %d : ", i+1)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[i] = klubA
		} else if skorB > skorA {
			pemenang[i] = klubB
		} else {
			pemenang[i] = "Draw"
		}perpus

		i++
	}

	for j := 0; j < i; j++ {
		fmt.Printf("Hasil %d : %s\n", j+1, pemenang[j])
	}

	fmt.Println("Pertandingan selesai")
}