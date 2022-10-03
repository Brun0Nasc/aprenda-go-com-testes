package arraysslices

func Soma(numeros [5]int) (soma int) {

	for _, num := range numeros {
		soma += num
	}
	return
}
