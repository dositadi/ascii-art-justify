package asciijustify

import (
	"strings"
)

func (j *Justify) GetSpacesIndex(s string) []int {
	var output []int

	offset := 0

	for {
		if offset > len(s)-1 {
			break
		}

		idx := strings.Index(s[offset:], " ")

		if idx == -1 {
			break
		}

		mainIndex := offset + idx

		output = append(output, mainIndex)

		offset = 1 + mainIndex
	}

	return output
}
