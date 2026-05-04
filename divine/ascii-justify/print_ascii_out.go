package asciijustify

import (
	"fmt"
	"slices"
	"strings"
)

func (j *Justify) PrintAscii(input [][][]string, width int) {
	var asciiArt strings.Builder

	for _, words := range input {

		if slices.Compare(words[0], []string{""}) == 0 {
			asciiArt.WriteString("\n")
			continue
		}

		for i := 0; i < 8; i++ {

			for k, word := range words {
				if slices.Compare(word, []string{" "}) == 0 {
					asciiArt.WriteString(strings.Repeat(" ", 40))
					continue
				}
				asciiArt.WriteString(words[k][i])
			}
			asciiArt.WriteString("\n")
		}
	}

	fmt.Print(asciiArt.String())
}

func (j *Justify) calcAmountOfSpace(terminalWidth int, words [][]string) int {
	totalGaps := len(words) - 1
	fmt.Println("TG: ", totalGaps)

	totalWordChars := 0

	for _, word := range words {
		for _, chars := range word {
			totalWordChars += len(chars)
		}
	}

	fmt.Println("TW: ", totalWordChars)

	availableSpace := terminalWidth - totalWordChars

	if totalGaps == 0 {
		return availableSpace
	}

	spaceAmount := availableSpace / totalGaps

	return spaceAmount
}
