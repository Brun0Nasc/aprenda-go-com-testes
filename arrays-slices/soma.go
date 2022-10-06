package arraysslices

func Soma(numeros []int) (soma int) {

	for _, num := range numeros {
		soma += num
	}
	return
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int //* Criando slice vazio sem tamanho definido

	for _, numeros := range numerosParaSomar {
		final := numeros[1:]
		somas = append(somas, Soma(final)) //* Percorrendo parâmetros passados e dando append do resultado da somas deles em somas  
	}

	return somas
}
