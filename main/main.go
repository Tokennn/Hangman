package main

import (
	"fmt"
	hangman "hang"
	"math/rand"
	"os"
	"time"
)

func main() {

	var data hangman.HangManData
	pointerData := &data
	randomWord, err := hangman.Randomly(os.Args[1])
	data.Tofind = randomWord
	data.Attempts = 10
	if err != nil {
		fmt.Println("Error:", err)
		// } else {
		// fmt.Printf("Random word : %s\n", data.Word)
	}

	for range data.Tofind {
		data.Word += "_"
	}

	// Révéler n lettres aléatoires dans le mot
	n := len(data.Tofind)/2 - 1
	revealed := make([]bool, len(data.Tofind))
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(data.Tofind))
		if revealed[randomIndex] {
			i--
		} else {
			revealed[randomIndex] = true
		}
	}
	newWord := ""
	for a, i := range revealed {
		if i {
			newWord += string(data.Tofind[a])
		} else {
			newWord += "_"
		}
	}
	data.Word = newWord
	fmt.Println(data.Tofind)

	// Afficher le mot avec les lettres révélées
	for i, c := range data.Tofind {
		if revealed[i] {
			fmt.Printf("%c ", c)
		} else {
			fmt.Print("_ ")
		}
	}

	fmt.Println()
	hangman.RechercheMot(pointerData)
	hangman.Validallletter(revealed)
	hangman.Partial(data.Word, revealed)
	hangman.Putwordinside("a", 0, "a")
}
