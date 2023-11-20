package hangman

func Validallletter(lettresTrouvees []bool) bool {
	for _, trouvee := range lettresTrouvees {
		if !trouvee {
			return false
		}
	}
	return true
}
