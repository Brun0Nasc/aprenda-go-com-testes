package arraysslices

func Soma(numeros []int) (soma int) {

	for _, num := range numeros {
		soma += num
	}
	return
}
