package iteracao

import (
	"testing"
	"fmt"
)

func TestRepetir(t *testing.T) {
	resultado := Repetir("a")
	esperado := "aaaaa"

	if resultado != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, resultado)
	}
}

func ExampleRepetir(){
	resultado := Repetir("a")
	fmt.Println(resultado)
	// Output: aaaaa
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}