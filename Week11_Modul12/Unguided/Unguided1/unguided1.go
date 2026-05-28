package main

import "fmt"

func main() {
	var pilihan int
	var calon [20]int

	for i := 0; i < 20; i++ {
		calon[i] = i + 1
	}

	var suara [21]int
	totalValid := 0

	for {
		fmt.Scan(&pilihan)
		if pilihan == 0 {
			break
		}

		valid := false
		for i := 0; i < 20; i++ {
			if pilihan == calon[i] {
				valid = true
			}
		}

		if valid {
			suara[pilihan]++
			totalValid++
		}
	}

	fmt.Println(totalValid)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("Calon %d: %d suara\n", i, suara[i])
		}
	}
}