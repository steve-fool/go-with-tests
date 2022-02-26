package main

import "testing"

func TestOla(t *testing.T) {

	verificarMensagemCorreta := func(t *testing.T, resultado string, esperado string){
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T){
		resultado := Ola("Chris", "portugues")
		esperado := "Olá, Chris"

		verificarMensagemCorreta(t, esperado, resultado)
	})

	t.Run("diz 'Olá, Mundo' quando uma string vazia for passada", func(t *testing.T){
		resultado := Ola("", "portugues")
		esperado := "Olá, Mundo"

		verificarMensagemCorreta(t, esperado, resultado)
	})

	t.Run("em espanhol", func(t *testing.T){
		resultado := Ola("Elodie", "espanhol")
		esperado := "Hola, Elodie"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em françês", func(t *testing.T){
		resultado := Ola("Marjorie", "françês")
		esperado := "Bonjour, Marjorie"
		verificarMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em dinamarques", func(t *testing.T){
		resultado := Ola("Lars", "dinamarques")
		esperado := "Hej, Lars"
		verificarMensagemCorreta(t, resultado, esperado)
	})
}