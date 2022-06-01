package exercise_2

//TODO: Доп задание

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func repeatSymbol(symbol rune, n int) []rune {
	var result []rune
	for i := 0; i < n; i++ {
		result = append(result, symbol)
	}

	return result
}

func primitiveExtract(str string) string {
	switch len(str) {
	case 0:
		return ""
	case 1:
		return str
	default:
		var result []rune
		symbols := []rune(str)
		for len(symbols) > 0 {
			if len(symbols) == 1 {
				if isDigit(symbols[0]) {
					return "некорректная строка"
				}
				result = append(result, symbols[0])
				break
			}
			fSymbol, sSymbol := symbols[0], symbols[1]
			if isDigit(fSymbol) {
				return "некорректная строка"
			}
			if isDigit(sSymbol) {
				result = append(result, repeatSymbol(fSymbol, int(sSymbol-'0'))...)
				symbols = symbols[2:]
			} else {
				result = append(result, fSymbol)
				symbols = symbols[1:]
			}
		}
		return string(result)
	}
}
