package main

import "fmt"


const( 
	espanhol = "espanhol"
	frances = "françês"
	dinamarques = "dinamarques"
)

const( 
	prefixoOlaPortuges = "Olá, "
	prefixoOlaEspanhol = "Hola, "
	prefixoOlaFrances = "Bonjour, "
	prefixoOlaDinamarques = "Hej, "
)

func Ola(nome string, idioma string) string {
	if nome == "" { nome = "Mundo" }

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string){
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case dinamarques:
		prefixo = prefixoOlaDinamarques
	default:
		prefixo = prefixoOlaPortuges
	}
	return
}

func main(){
	fmt.Println(Ola("", ""))
}