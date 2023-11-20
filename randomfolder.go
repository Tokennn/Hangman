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
		return "", fmt.Errorf("we can't open the folder %s : %v", filename, err)
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
		return "", fmt.Errorf("error while reading file %s : %v", filename, err)
	}

	if len(words) == 0 {
		return "", fmt.Errorf("the folder %s is empty", filename)
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	return words[randomIndex], nil
}
