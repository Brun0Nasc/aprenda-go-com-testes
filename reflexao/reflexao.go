package reflexao

import "reflect"

func percorre(x interface{}, fn func(entrada string)){
	valor := reflect.ValueOf(x) //* valor de
	
	for i := 0; i < valor.NumField(); i++ {
		campo := valor.Field(i)
		fn(campo.String())
	}
}