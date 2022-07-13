package figuras

import "fmt"

type Cuadrado struct {
	Anc float64
	Alt float64
}

func (p *Cuadrado) Area() float64 {
	fmt.Println("area del cuadrado")
	return p.Anc * p.Alt
}

func (p *Cuadrado) Perimetro() float64 {
	fmt.Println("perimetro del cuadrado")
	return 2*p.Anc + 2*p.Alt
}

func InsertC() {
	cuadrado := Cuadrado{
		Anc: 2,
		Alt: 3,
	}
	CalculoA(&cuadrado)
}
