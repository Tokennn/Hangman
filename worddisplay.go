package hangman

import (
	"math/rand"
	"strings"
)

func Displaywords(words string) string {
	numbreveal := (len(words) / 2) - 1
	tabword := make([]string, len(words))
	for i := range tabword {
		tabword[i] = "_"
	}
	for i := 0; i < numbreveal; i++ {
		max := len(words) - 1
		indexrandomword := rand.Intn(max)
		randomletter := string(words[indexrandomword])
		for j := 0; j < len(words); j++ {
			if string(words[j]) == randomletter {
				tabword[j] = randomletter
			}
		}
	}
	return strings.Join(tabword, "")
}
