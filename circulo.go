package figuras

import (
	"fmt"
	"math"
)

type Circulo struct {
	Rad float64
}

func (p *Circulo) Area() float64 {
	fmt.Println("area del circulo")
	return math.Pi * (p.Rad * p.Rad)
}

func (p *Circulo) Perimetro() float64 {
	fmt.Println("perimetro del circulo")
	return 2 * math.Pi * p.Rad
}

func InsertCir() {
	circulo := Circulo{
		Rad: 2,
	}
	CalculoA(&circulo)

}
