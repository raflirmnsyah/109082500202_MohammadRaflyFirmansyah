package main

import "fmt"

func main() {
	var pilihan int
	var suara [21]int
	suaraMasuk := 0
	suaraSah := 0
	for {
		fmt.Scan(&pilihan)
		if pilihan == 0 {
			break
		}
		suaraMasuk++
		valid := false
		for i := 1; i <= 20; i++ {
			if pilihan == i {
				valid = true
			}
		}
		if valid {
			suara[pilihan]++
			suaraSah++
		}
	}
	ketua := 1
	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			ketua = i
		}
	}
	wakil := 1
	if ketua == 1 {
		wakil = 2
	}
	for i := 1; i <= 20; i++ {
		if i != ketua {
			if suara[i] > suara[wakil] {
				wakil = i
			}
		}
	}
	fmt.Println("Suara Masuk :", suaraMasuk)
	fmt.Println("Suara Sah   :", suaraSah)
	fmt.Println("Ketua RT    :", ketua)
	fmt.Println("Wakil RT    :", wakil)
}