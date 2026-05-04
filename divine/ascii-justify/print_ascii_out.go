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

		spaces := j.calcAmountOfSpace(width, words)

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

	totalWordChars := 0

	for _, word := range words {
		totalWordChars += len(slices.Max(word))
	}

	availableSpace := terminalWidth - totalWordChars

	if availableSpace < 0 {
		return 4
	}

	if totalGaps == 0 {
		return availableSpace
	}

	spaceAmount := availableSpace / totalGaps

	return spaceAmount
}
