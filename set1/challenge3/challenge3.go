package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/agnivade/levenshtein"
)

func main() {
	input, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		log.Fatal(err)
	}
	output := decryptHex(input)
	fmt.Println(output)
}

func decryptHex(input []byte) string {
	plaintext := make([]byte, len(input))
	scores := make(map[string]float32)

	// Guess at the cipher by iterating over characters
	for i := 48; i <= 135; i++ {
		for j := 0; j < len(input); j++ {
			plaintext[j] = input[j] ^ byte(rune(i)) // Just repeat the same character as the ciphertext
		}
		scores[string(plaintext)] = scorePlaintext(string(plaintext))
	}

	highestScore := float32(0)
	highestScoreText := ""
	for text, score := range scores {
		if score > highestScore {
			highestScore = score
			highestScoreText = text
		}
	}

	return highestScoreText
}

// Generates a score for how likely a plaintext is readable.
// Returns a result between 0 and 1.
func scorePlaintext(input string) float32 {
	score := float32(0)

	// Try to find words
	words := strings.Split(strings.ToLower(input), " ")
	for _, word := range words {
		hasVowels := strings.ContainsAny(word, "aeiou")
		hasConsonants := strings.ContainsAny(word, "bcdfghjklmnpqrstvwxyz")
		lengthIsWordlike := len(word) < 8
		if hasVowels && hasConsonants && lengthIsWordlike {
			score++
		}
	}

	// Analyze character frequency
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	englishCharacterFrequency := "ETAONRISHDLFCMUGYPWBVKJXZQ"

	// Initialize frequency
	inputCharacterFrequencyMap := make(map[rune]int, 26) // Holds the data
	inputCharacterFrequency := []rune(alphabet)          // To be sorted

	for _, character := range alphabet {
		inputCharacterFrequencyMap[character] = 0
	}
	for _, character := range strings.ToUpper(input) {
		if strings.ContainsRune(alphabet, character) {
			inputCharacterFrequencyMap[character]++
		}
	}

	// https://code-maven.com/slides/golang/solution-count-characters-sort-by-frequency
	sort.Slice(inputCharacterFrequency, func(i int, j int) bool {
		return inputCharacterFrequencyMap[inputCharacterFrequency[i]] > inputCharacterFrequencyMap[inputCharacterFrequency[j]]
	})

	// Compute Levenshtein distance
	editDistance := levenshtein.ComputeDistance(englishCharacterFrequency, string(inputCharacterFrequency))

	// Get a percentage. I'm less concerned with the math itself - i.e. an equal distribution from 0-1, etc - but as long as we
	// can differentiate from something that looks like words and garbage, it will do.
	score = ((score / float32(len(words))) + (1 - float32(editDistance)/26)) / 2
	return score
}
