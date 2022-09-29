package docstrings

import (
	"fmt"
	"testing"
)

func TestComparar(t *testing.T) {
	comparar := Comparar("a", "b")
	esperado := -1

	if comparar != esperado {
		t.Errorf("resultado '%d', esperado '%d'", comparar, esperado)
	}
}

func ExampleComparar(){
	comparar := Comparar("a", "a")
	fmt.Println(comparar)

	// Output: 0
}

func BenchmarkComparar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Comparar("a", "b")
	}
}