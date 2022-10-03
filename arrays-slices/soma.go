package arraysslices

func Soma(numeros []int) (soma int) {

	for _, num := range numeros {
		soma += num
	}
	return
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	quantidadeDeNumeros := len(numerosParaSomar)
	somas = make([]int, quantidadeDeNumeros) //* O make permite criar um slice com uma capacidade inicial de *len* *numerosParaSomar* que precisamos percorrer 

	for i, numeros := range numerosParaSomar {
		somas[i] = Soma(numeros)
	}

	return
}
