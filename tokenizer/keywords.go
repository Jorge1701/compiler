package tokenizer

var keywords = []string{
	"salir",
}

func isKeyWord(s string) bool {
	for _, kw := range keywords {
		if s == kw {
			return true
		}
	}

	return false
}
