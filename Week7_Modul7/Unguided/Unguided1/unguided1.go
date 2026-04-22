package main
import "fmt"

type Suhu struct {
	celcius float64
}

func (s Suhu) Reamur() float64 {
	return (4.0 / 5.0) * s.celcius
}

func (s Suhu) Fahrenheit() float64 {
	return (9.0/5.0)*s.celcius + 32
}

func (s Suhu) Kelvin() float64 {
	return s.celcius + 273.15
}

func main() {
	var input float64
	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius): ")
	fmt.Scan(&input)
	s := Suhu{celcius: input}
	fmt.Printf("\n%.0f celcius = %.1f reamur\n", s.celcius, s.Reamur())
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", s.celcius, s.Fahrenheit())
	fmt.Printf("%.0f celcius = %.2f kelvin\n", s.celcius, s.Kelvin())
}