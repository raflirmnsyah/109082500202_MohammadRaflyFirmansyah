package main

import "fmt"

const NMAX = 100000

type partai struct {
	nama  int
	suara int
}

type tabPartai [NMAX]partai

func posisi(t tabPartai, n int, x int) int {
	for i := 0; i < n; i++ {
		if t[i].nama == x {
			return i
		}
	}
	return -1
}

func main() {
	var p tabPartai
	var x int
	var n int = 0
	var idx int

	fmt.Scan(&x)

	for x != -1 {
		idx = posisi(p, n, x)

		if idx == -1 {
			p[n].nama = x
			p[n].suara = 1
			n++
		} else {
			p[idx].suara++
		}

		fmt.Scan(&x)
	}

	for pass := 1; pass < n; pass++ {
		temp := p[pass]
		i := pass - 1

		for i >= 0 && p[i].suara < temp.suara {
			p[i+1] = p[i]
			i--
		}

		p[i+1] = temp
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d(%d)", p[i].nama, p[i].suara)

		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}