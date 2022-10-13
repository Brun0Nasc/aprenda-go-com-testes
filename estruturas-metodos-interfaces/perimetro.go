package estruturasmetodosinterfaces

import "math"

type Forma interface { //* Interface que recebe uma função Area do tipo float64
	Area() float64
}

type Retangulo struct { //* Estrutura de retangulo que recebe largura e altura
	Largura float64
	Altura  float64
}

type Circulo struct { //* Estrutura de círculo que recebe raio
	Raio float64
}

type Triangulo struct { //* Estrutura de triangulo que recebe base e altura
	Base   float64
	Altura float64
}

func (r Retangulo) Perimetro() float64 { //* Método para calcular o perímetro de um retangulo
	return 2 * (r.Largura + r.Altura)
}

func (r Retangulo) Area() float64 { //* Método para calcular a area do retangulo, a função área desse método é o que a interface requer
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 { //* Método para calcular area do circulo, nos mesmos moldes do método de retângulo
	return math.Pi * c.Raio * c.Raio
}

func (t Triangulo) Area() float64 { //* Método para calcular a area do triangulo
	return (t.Base * t.Altura) * 0.5
}
