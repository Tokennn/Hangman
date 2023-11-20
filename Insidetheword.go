package hangman

func Putwordinside(motUniversel string, index int, lettre string) string {
	lettres := []rune(motUniversel)
	lettres[index] = []rune(lettre)[0]
	return string(lettres)
}
