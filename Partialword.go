package hangman

func Partial(mot string, lettresTrouvees []bool) string {
	partiel := ""

	for i, char := range mot {
		if lettresTrouvees[i] {
			partiel += string(char)
		} else {
			partiel += "_"
		}
	}

	return partiel
}
