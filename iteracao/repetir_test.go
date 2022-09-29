package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 7)
	esperado := "aaaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s', mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 7)
	}
}

func ExampleRepetir() {
	repeticoes := Repetir("a", 4)
	fmt.Println(repeticoes)

	// Output: aaaa
}
