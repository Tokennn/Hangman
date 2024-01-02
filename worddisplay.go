package hangman

import (
	"math/rand"
	"strings"
)

//	func DisplayWord(revealed []bool, word string) {
//		for i, c := range word {
//			if revealed[i] {
//				fmt.Printf("%c ", c)
//			} else {
//				fmt.Print("_ ")
//			}
//		}
//	}
func Displayword(words string) string {
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
