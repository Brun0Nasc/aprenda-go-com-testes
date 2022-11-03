package reflexao

import "reflect"

func percorre(x interface{}, fn func(entrada string)) {
	valor := obtemValor(x)

	quantidadeDeValores := 0
	var obtemCampo func(int) reflect.Value

	switch valor.Kind(){
	case reflect.String:
		fn(valor.String())
	case reflect.Struct:
		quantidadeDeValores = valor.NumField()
		obtemCampo = valor.Field
	case reflect.Slice, reflect.Array:
		quantidadeDeValores = valor.Len()
		obtemCampo = valor.Index
	case reflect.Map:
		for _, chave := range valor.MapKeys() {
			percorre(valor.MapIndex(chave).Interface(), fn)
		}
	}

	for i := 0; i < quantidadeDeValores; i++ {
		percorre(obtemCampo(i).Interface(), fn)
	}

	// if valor.Kind() == reflect.Slice {
	// 	for i := 0; i < valor.Len(); i++ {
	// 		percorre(valor.Index(i).Interface(), fn)
	// 	}
	// 	return
	// }

	// for i := 0; i < valor.NumField(); i++ {
	// 	campo := valor.Field(i)

	// 	switch campo.Kind() {
	// 	case reflect.String:
	// 		fn(campo.String())
	// 	case reflect.Struct:
	// 		percorre(campo.Interface(), fn)
	// 	}

	// }
}

func obtemValor(x interface{}) reflect.Value {
	valor := reflect.ValueOf(x)

	if valor.Kind() == reflect.Ptr {
		valor = valor.Elem()
	}

	return valor
}
