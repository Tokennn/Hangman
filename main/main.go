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
	randomWord, err := hangman.Randomly(os.Args[1])
	data.Word = randomWord

	if err != nil {
		fmt.Println("Erreur:", err)
	} else {
		fmt.Printf("Mot aléatoire : %s\n", data.Word)
	}

	// Révéler n lettres aléatoires dans le mot
	n := len(data.Word)/2 - 1
	revealed := make([]bool, len(data.Word))
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(data.Word))
		if !revealed[randomIndex] {
			revealed[randomIndex] = true
		} else {
			i--
		}
	}

	// Afficher le mot avec les lettres révélées
	for i, c := range data.Word {
		if revealed[i] {
			fmt.Printf("%c ", c)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
	hangman.Reveal(data.Word, data.Word, "a")
}
