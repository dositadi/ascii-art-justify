package asciijustify

import (
	"fmt"
	"slices"
	"strings"
)

func (j *Justify) PrintAscii(input [][][]string, width int) {
	var asciiArt strings.Builder

	fmt.Println("input length: ", len(input))

	for _, words := range input {

		if slices.Compare(words[0], []string{""}) == 0 {
			asciiArt.WriteString("\n")
			continue
		}

		spaces := j.calcAmountOfSpace(width, words)
		fmt.Println("Spaces: ", spaces)

		for i := 0; i < 8; i++ {

			for k, word := range words {
				if slices.Compare(word, []string{" "}) == 0 {
					asciiArt.WriteString(strings.Repeat(" ", spaces))
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
	totalGaps := 0

	for _, word := range words {
		if slices.Compare(word, []string{" "}) == 0 {
			totalGaps += 1
		}
	}

	fmt.Println("TG: ", totalGaps)

	totalWordChars := 0

	for _, word := range words {
		totalWordChars += len(slices.Max(word))
		fmt.Println("maximal word: ", slices.Max(word))
	}

	fmt.Println("TW: ", totalWordChars)

	availableSpace := terminalWidth - totalWordChars
	fmt.Println("AS: ", availableSpace)

	if availableSpace < 0 {
		return 4
	}

	if totalGaps == 0 {
		return availableSpace
	}

	spaceAmount := availableSpace / totalGaps

	return spaceAmount
}
