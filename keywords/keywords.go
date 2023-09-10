package keywords

type KeyWord string

const (
	SALIR KeyWord = "salir"
)

var keywords = []KeyWord{
	SALIR,
}

func IsKeyWord(s string) bool {
	for _, kw := range keywords {
		if s == string(kw) {
			return true
		}
	}

	return false
}
