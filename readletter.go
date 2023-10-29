package hangman

import (
	"fmt"
	"os"
)

func ReadLetter(data *HangManData) { //affiche une lettre
	if len(os.Args) < 2 {
		fmt.Println("Veuillez fournir un mot en argument.")
		os.Exit(1)
	}

	mot := os.Args[1]

	lettre := genererLettre(mot)

	fmt.Println(lettre)
}
