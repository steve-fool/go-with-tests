package colecoes

import (
	"testing"
	"reflect"
)

func TestSoma(t *testing.T) {
	t.Run("coleção com qualquer tamanho", func(t *testing.T) {
		numeros := []int{1,2,3}
		
		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", resultado, esperado, numeros)
		}
	})
}

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1,2}, []int{0,9})
	esperado := []int{3,9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {

	verificaSomar := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		if !reflect.DeepEqual(resultado, esperado){
			t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
		}
	}

	t.Run("realiza a soma de alguns slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1,2}, []int{0,9})
		esperado := []int{2,9}

		verificaSomar(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3,4,5})
		esperado := []int{0,9}
		
		verificaSomar(t, resultado, esperado)
	})
}