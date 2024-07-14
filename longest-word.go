package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func longest(sentence string) (string, int) {
	words := strings.Fields(sentence)
	var longestWord string
	maxLength := 0

	for _, word := range words {
		if len(word) > maxLength {
			longestWord = word
			maxLength = len(word)
		}
	}

	return longestWord, maxLength
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukan kalimat: ")
	sentence, _ := reader.ReadString('\n')

	word, length := longest(sentence)
	fmt.Printf("kata terpanjang adalah = %s: %d character\n", word, length)
}
