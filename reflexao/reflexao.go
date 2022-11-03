package reflexao

import "reflect"

func percorre(x interface{}, fn func(entrada string)){
	valor := reflect.ValueOf(x) //* valor de
	campo := valor.Field(0) //* campo
	fn(campo.String())
}