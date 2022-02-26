package inteiros

import (
	"testing"
	"fmt"
)

func TestAdiciona(t *testing.T){
	soma := Adiciona(1, 2)
	esperado := 3

	if soma != esperado {
		t.Errorf("esperado '%d', esperado '%d'", soma, esperado)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(5, 7)
	fmt.Println(soma)
	// Output: 12
}