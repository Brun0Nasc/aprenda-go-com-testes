package reflexao

import (
	"reflect"
	"testing"
)

func TestPercorre(t *testing.T) {

	casos := []struct {
		Nome              string
		Entrada           interface{}
		ChamadasEsperadas []string
	}{
		{
			"Structs com um campo string",
			struct {
				Nome string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct com dois campos do tipo string",
			struct {
				Nome string
				Cidade string
			}{"Chris", "Londres"},
			[]string{"Chris", "Londres"},
		},
	}

	for _, teste := range casos {
		t.Run(teste.Nome, func(t *testing.T) {
			var resultado []string
			percorre(teste.Entrada, func(entrada string) {
				resultado = append(resultado, entrada)
			})

			if !reflect.DeepEqual(resultado, teste.ChamadasEsperadas){
				t.Errorf("resultado '%v', esperado '%v'", resultado, teste.ChamadasEsperadas)
			}
		})
	}

	// esperado := "Chris"
	// var resultado []string

	// x := struct {
	// 	Nome string
	// }{esperado}

	// percorre(x, func(entrada string) {
	// 	resultado = append(resultado, entrada)
	// })

	// if len(resultado) != 1 {
	// 	t.Errorf("número incorreto de chamadas de função: resultado %d, esperado %d", len(resultado), 1)
	// }

	// if resultado[0] != esperado {
	// 	t.Errorf("resultado '%v', esperado '%v'", resultado[0], esperado)
	// }
}