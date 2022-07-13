package figuras

import "fmt"

type Calculo interface {
	Area() float64
	Perimetro() float64
}

func CalculoA(al Calculo) {
	fmt.Println(al.Area())
	fmt.Println(al.Perimetro())
}
