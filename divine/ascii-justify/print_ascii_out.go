package asciijustify

import (
	"fmt"
	"slices"
	"strings"
)

func (j *Justify) PrintAscii(input [][][]string) {
	var asciiArt strings.Builder

	for _, word := range input {
		if slices.Compare(word[0], []string{""}) == 0 {
			fmt.Println("Entered")
			asciiArt.WriteString("\n")
			continue
		}

		for j := 0; j < 8; j++ {
			for k := range len(word) {
				asciiArt.WriteString(word[k][j])
			}
			asciiArt.WriteString("\n")
		}
	}

	fmt.Print(asciiArt.String())
}
