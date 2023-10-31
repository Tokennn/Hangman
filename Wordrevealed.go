package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Reveal(mot, motCache, lettre string) string {

	// 	revele := []byte(motCache)
	// 	for i := range mot {
	// 		if mot[i:i+1] == lettre {
	// 			revele[i] = mot[i]
	// 		}
	// 	}
	// 	return string(revele)
	// }

	// word := make([]string, 0)
	found := false

	for !found {
		fmt.Print("Entrez une lettre : ")
		reader := bufio.NewReader(os.Stdin)
		letter, _ := reader.ReadString('\n')
		letter = strings.TrimSpace(letter)

		if strings.Contains(mot, letter) {
			fmt.Println("Lettre trouvée !")
		} else {
			fmt.Println("Lettre non trouvée.")
		}

		found = strings.Contains(mot, letter)
	}

	fmt.Println("Mot trouvé !")
	return mot
}
