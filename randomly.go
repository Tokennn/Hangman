package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Randomly(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("impossible d'ouvrir le fichier %s : %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	words := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		wordList := strings.Fields(line)
		words = append(words, wordList...)
	}

	if err := scanner.Err(); err != nil { //
		return "", fmt.Errorf("erreur lors de la lecture du fichier %s : %v", filename, err)
	}

	if len(words) == 0 {
		return "", fmt.Errorf("le fichier %s est vide", filename)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	fmt.Println("test : " + words[randomIndex])
	return words[randomIndex], nil
}
