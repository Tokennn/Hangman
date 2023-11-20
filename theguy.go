package hangman

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Showhang(w *HangManData) {
	pendue, err := os.Open("hangman.txt")
	var tab []string

	if err != nil {
		log.Fatal(err)
	}
	defer pendue.Close()

	scanner := bufio.NewScanner(pendue)

	for scanner.Scan() {
		tab = append(tab, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if w.Attempts == 9 {

		index := 6
		pendue1 := tab[index]
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println(pendue1)

	}

	if w.Attempts == 8 {

		fmt.Println()
		for i := 9; i <= 14; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}

	if w.Attempts == 7 {

		fmt.Println()
		for i := 16; i <= 22; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if w.Attempts == 6 {

		fmt.Println()
		for i := 24; i <= 30; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if w.Attempts == 5 {

		fmt.Println()
		for i := 32; i <= 38; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if w.Attempts == 4 {

		fmt.Println()
		for i := 40; i <= 46; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if w.Attempts == 3 {

		fmt.Println()
		for i := 48; i <= 54; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if w.Attempts == 2 {

		fmt.Println()
		for i := 56; i <= 62; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if w.Attempts == 1 {

		fmt.Println()
		for i := 64; i <= 70; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}
	if w.Attempts == 0 {

		fmt.Println()
		for i := 72; i <= 78; i++ {
			pendue1 := tab[i]
			fmt.Println(pendue1)
		}

	}

}
