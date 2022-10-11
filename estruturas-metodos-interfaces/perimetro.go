package estruturasmetodosinterfaces

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura float64
}

type Circulo struct {
	Raio float64
}

func (r Retangulo) Perimetro() float64 {
	return 2 * (r.Largura + r.Altura)
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}