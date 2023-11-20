package hangman

import "fmt"

func DisplayWord(revealed []bool, word string) {
	for i, c := range word {
		if revealed[i] {
			fmt.Printf("%c ", c)
		} else {
			fmt.Print("_ ")
		}
	}
}
