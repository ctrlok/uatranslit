package uatranslit

// Letter is an atomic struct for letters.
type Letter struct {
	PositionFirst []rune
	PositionOther []rune
}

// Letters is an array of letter transliteration.
var Letters = map[rune]Letter{
	'А': Letter{PositionOther: []rune{'A'}},
	'а': Letter{PositionOther: []rune{'a'}},
	'Б': Letter{PositionOther: []rune{'B'}},
	'б': Letter{PositionOther: []rune{'b'}},
	'В': Letter{PositionOther: []rune{'V'}},
	'в': Letter{PositionOther: []rune{'v'}},
	'Г': Letter{PositionOther: []rune{'H'}},
	'г': Letter{PositionOther: []rune{'h'}},
	'Ґ': Letter{PositionOther: []rune{'G'}},
	'ґ': Letter{PositionOther: []rune{'g'}},
	'Д': Letter{PositionOther: []rune{'D'}},
	'д': Letter{PositionOther: []rune{'d'}},
	'Е': Letter{PositionOther: []rune{'E'}},
	'е': Letter{PositionOther: []rune{'e'}},
	'Є': Letter{PositionOther: []rune{'I', 'e'}, PositionFirst: []rune{'Y', 'e'}},
	'є': Letter{PositionOther: []rune{'i', 'e'}, PositionFirst: []rune{'y', 'e'}},
	'Ж': Letter{PositionOther: []rune{'Z', 'h'}},
	'ж': Letter{PositionOther: []rune{'z', 'h'}},
	'З': Letter{PositionOther: []rune{'Z'}},
	'з': Letter{PositionOther: []rune{'z'}},
	'И': Letter{PositionOther: []rune{'Y'}},
	'и': Letter{PositionOther: []rune{'y'}},
	'І': Letter{PositionOther: []rune{'I'}},
	'і': Letter{PositionOther: []rune{'i'}},
	'Ї': Letter{PositionOther: []rune{'I'}, PositionFirst: []rune{'Y', 'i'}},
	'ї': Letter{PositionOther: []rune{'i'}, PositionFirst: []rune{'y', 'i'}},
	'Й': Letter{PositionOther: []rune{'I'}, PositionFirst: []rune{'Y'}},
	'й': Letter{PositionOther: []rune{'i'}, PositionFirst: []rune{'y'}},
	'К': Letter{PositionOther: []rune{'K'}},
	'к': Letter{PositionOther: []rune{'k'}},
	'Л': Letter{PositionOther: []rune{'L'}},
	'л': Letter{PositionOther: []rune{'l'}},
	'М': Letter{PositionOther: []rune{'M'}},
	'м': Letter{PositionOther: []rune{'m'}},
	'Н': Letter{PositionOther: []rune{'N'}},
	'н': Letter{PositionOther: []rune{'n'}},
	'О': Letter{PositionOther: []rune{'O'}},
	'о': Letter{PositionOther: []rune{'o'}},
	'П': Letter{PositionOther: []rune{'P'}},
	'п': Letter{PositionOther: []rune{'p'}},
	'Р': Letter{PositionOther: []rune{'R'}},
	'р': Letter{PositionOther: []rune{'r'}},
	'С': Letter{PositionOther: []rune{'S'}},
	'с': Letter{PositionOther: []rune{'s'}},
	'Т': Letter{PositionOther: []rune{'T'}},
	'т': Letter{PositionOther: []rune{'t'}},
	'У': Letter{PositionOther: []rune{'U'}},
	'у': Letter{PositionOther: []rune{'u'}},
	'Ф': Letter{PositionOther: []rune{'F'}},
	'ф': Letter{PositionOther: []rune{'f'}},
	'Х': Letter{PositionOther: []rune{'K', 'h'}},
	'х': Letter{PositionOther: []rune{'k', 'h'}},
	'Ц': Letter{PositionOther: []rune{'T', 's'}},
	'ц': Letter{PositionOther: []rune{'t', 's'}},
	'Ч': Letter{PositionOther: []rune{'C', 'h'}},
	'ч': Letter{PositionOther: []rune{'c', 'h'}},
	'Ш': Letter{PositionOther: []rune{'S', 'h'}},
	'ш': Letter{PositionOther: []rune{'s', 'h'}},
	'Щ': Letter{PositionOther: []rune{'S', 'h', 'c', 'h'}},
	'щ': Letter{PositionOther: []rune{'s', 'h', 'c', 'h'}},
	'Ю': Letter{PositionOther: []rune{'I', 'u'}, PositionFirst: []rune{'Y', 'u'}},
	'ю': Letter{PositionOther: []rune{'i', 'u'}, PositionFirst: []rune{'y', 'u'}},
	'Я': Letter{PositionOther: []rune{'I', 'a'}, PositionFirst: []rune{'Y', 'a'}},
	'я': Letter{PositionOther: []rune{'i', 'a'}, PositionFirst: []rune{'y', 'a'}},
}

// ReplaceUARune is a func which replace UA rune with list of ENG runes if UA rune exist
// if UA rune dosen't exist - return initial rune in a list
func ReplaceUARune(uaRune rune, isFist bool) []rune {
	letter, exist := Letters[uaRune]
	if !exist {
		return []rune{uaRune}
	}
	if isFist && letter.PositionFirst != nil {
		return letter.PositionFirst
	}
	return letter.PositionOther
}

// ReplaceUARunes is a func which replace list of UA runes with list of ENG runes if UA runes exist
// if UA rune dosen't exist - return initial rune in a list
func ReplaceUARunes(uaRunes []rune) (engRunes []rune) {
	isFirst := true
	for _, rune := range uaRunes {
		newRunes := ReplaceUARune(rune, isFirst)
		engRunes = append(engRunes, newRunes...)
		if newRunes[0] == rune {
			isFirst = true
			continue
		}
		isFirst = false
	}
	return engRunes
}
