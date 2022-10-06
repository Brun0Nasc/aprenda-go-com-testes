package arraysslices

func Soma(numeros []int) (soma int) {

	for _, num := range numeros {
		soma += num
	}
	return
}

func SomaTudo(numerosParaSomar ...[]int) []int {
	var somas []int //* Criando slice vazio sem tamanho definido

	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma(numeros)) //* Percorrendo parâmetros passados e dando append do resultado da somas deles em somas  
	}

	return somas
}
