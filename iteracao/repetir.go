package iteracao

// Repetir recebe uma string e um inteiro, o inteiro é a quatidade de vezes que a string vai repetir
func Repetir(s string, n int) (res string) {
	for i := 0; i < n; i ++ {
		res += s
	}
	return
}
