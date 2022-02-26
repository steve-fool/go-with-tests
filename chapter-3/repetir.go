package iteracao

const qtdrepeticoes = 5

func Repetir(value string) string {
	var repeticoes string

	for i := 0; i < qtdrepeticoes; i++ {
		repeticoes +=  value
	}
	return repeticoes
}