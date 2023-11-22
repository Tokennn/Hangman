package hangman

import (
	"fmt"
	"strings"
)

func Reveal(data *HangManData) {
	goal := data.Tofind
	// tentativesRestantes := 10
	gamecondit := true

	// pendu := make([]string, 0)

	for i := 0; i < 7; i++ {
		// pendu = append(pendu, "")
	}

	for gamecondit {
		input := ""
		// fmt.Print("Current word: ", "\n")
		// Partial(motUniversel, lettresTrouvees)
		fmt.Println("Attempts: ", data.Attempts)
		fmt.Print("Choose : ")
		fmt.Scan(&input)

		if len(input) > 1 {
			if input == data.Tofind {
				fmt.Println("Congrats :", data.Tofind)
				return
			} else {
				data.Attempts += -1
				fmt.Println("Wrong word. Try again.")
			}
		}

		input = strings.ToLower(input)
		trouvee := false

		for i, char := range goal {
			if string(char) == input {
				trouvee = true
				newWord := ""
				for x := range data.Word {
					if x == i {
						newWord += string(char)
					} else {
						newWord += string(data.Word[x])
					}
				}
				data.Word = newWord
			}
		}

		if trouvee {
			fmt.Println("Letter found!")
		} else {
			fmt.Println("Letter not found. Try again.")
			data.Attempts--
			// Showhang(pendu, 10-tentativesRestantes)
		}

		if data.Attempts == 0 {
			fmt.Println("You lose", data.Tofind)
			return
		}

		fmt.Println("Word: ", data.Word)

		if data.Tofind == data.Word {
			fmt.Println("Congrats :", data.Word)
			return
		}
		Showhang(data)
	}
}
