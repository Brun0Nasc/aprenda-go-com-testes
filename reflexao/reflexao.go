package reflexao

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	valor := reflect.ValueOf(x) //* valor de

	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)

		if campo.Kind() == reflect.String { // Tipo
			fn(campo.String())
		}

		if campo.Kind() == reflect.Struct {
			percorre(campo.Interface(), fn) //* Isso é recursividade :O
		}
	}
}
