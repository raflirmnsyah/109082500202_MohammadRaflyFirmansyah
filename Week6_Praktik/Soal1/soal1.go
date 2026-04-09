package main

import "fmt"

const pi float64 = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	selisih := m1 - m2
	fmt.Printf("\nMassa tabung kiri  : %.2f\n", m1)
	fmt.Printf("Massa tabung kanan : %.2f\n", m2)
	fmt.Printf("Selisih massa      : %.2f\n", selisih)
	if selisih == 0 {
		fmt.Println("Balance")
	} else if selisih > 0 {
		fmt.Println("Tabung kiri lebih berat")
	} else {
		fmt.Println("Tabung kanan lebih berat")
	}
}

func main() {
	var r float64
	var tkiri, tkanan float64
	var mjkiri, mjkanan float64
	var mkiri, mkanan float64
	fmt.Print("Masukkan jari-jari alas tabung: ")
	fmt.Scan(&r)
	fmt.Print("Masukkan tinggi zat cair tabung kiri: ")
	fmt.Scan(&tkiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri: ")
	fmt.Scan(&mjkiri)
	fmt.Print("Masukkan tinggi zat cair tabung kanan: ")
	fmt.Scan(&tkanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan: ")
	fmt.Scan(&mjkanan)
	mkiri = massa(r, tkiri, mjkiri)
	mkanan = massa(r, tkanan, mjkanan)
	display(mkiri, mkanan)
}