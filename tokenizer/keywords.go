package tokenizer

type KeyWord string

const (
	SALIR KeyWord = "salir"
)

var keywords = []KeyWord{
	SALIR,
}

func isKeyWord(s string) bool {
	for _, kw := range keywords {
		if s == string(kw) {
			return true
		}
	}

	return false
}
