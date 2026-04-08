package main

import "fmt"

func cetakbintang(n int, baris int) {
	if baris > n {
		return
	}
	for i := 0; i < baris; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	cetakbintang(n, baris+1)
}

func main() {
	var n int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)
	cetakbintang(n, 1)
}