package estruturasmetodosinterfaces

import "testing"

func TestAreaPerimetro(t *testing.T) { //* testando a função de calcular perímetro da forma normal

	retangulo := Retangulo{10.0, 10.0}
	resultado := retangulo.Perimetro()
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}

}

func TestArea(t *testing.T) { //* testando os metodos de calcular area das formas
	testesArea := []struct { //* criando uma struct para receber structs das formas
		nome string //* o campo nome vai servir para nomear os testes que serão feitos
		forma Forma //* forma vai receber a interface de forma que se refere a cada estrutura de forma declarada no outro arquivo
		temArea float64 //* campo para receber a área que espera-se obter no calculo da função area
	} {
		{nome: "Retângulo", forma: Retangulo{Largura: 12, Altura: 6}, temArea:72.0}, //* passando cada forma com nome, forma e area esperada
		{nome: "Círculo", forma: Circulo{Raio: 10}, temArea: 314.1592653589793},
		{nome: "Triângulo", forma: Triangulo{Base: 12, Altura: 6}, temArea: 36.0},
	}

	for _, tt := range testesArea { //* o for vai percorrer o testesArea para testar cada estrutura declarada unitariamente
		t.Run(tt.nome, func(t *testing.T) { //* o t.Run vai organizar melhor o teste, exibindo o nome de cada teste
			//* isso é útil pois se um arquivo de testes tiver muitos testes a serem feitos, o log vai aparecer nomeado e será mais fácil encontrar no código
			resultado := tt.forma.Area() //* tt é cada elemento único que está sendo iterado dentro de testesArea, forma é o campo que recebe Forma
			//* a partir do que é recebido nesse campo, estruturas singulares, é possível chamar as funções dos métodos declarados para cada forma
			if resultado != tt.temArea {
				t.Errorf("resultado %.2f esperado %.2f", resultado, tt.temArea)
			}	
		})

	}
}
