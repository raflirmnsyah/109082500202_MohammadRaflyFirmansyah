package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var lingkaran [2]Lingkaran
	var titik Titik
	
	for i := 0; i < 2; i++ {
		fmt.Scan(&lingkaran[i].pusat.x, &lingkaran[i].pusat.y, &lingkaran[i].r)
	}

	fmt.Scan(&titik.x, &titik.y)

	dalam1 := didalam(lingkaran[0], titik)
	dalam2 := didalam(lingkaran[1], titik)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}