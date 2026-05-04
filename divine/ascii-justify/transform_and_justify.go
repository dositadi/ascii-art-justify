package asciijustify

import (
	"fmt"
)

func (j *Justify) TransformAndJustify(input []string, font string) [][][]string {
	var output [][][]string

	/* slice := [][]string{{"abc"}, {"b"}, {"c"}, {"d"}}

	slice[2] = slices.Repeat(slice[2], 3)

	fmt.Println(slice) */

	for _, word := range input {
		var chars [][]string

		if word == "" {
			chars = append(chars, []string{""})
			output = append(output, chars)
			continue
		}

		fmt.Println(len(word))

		for _, char := range word {
			if char < ' ' || char > '~' {
				continue
			}

			ascii, err := j.ReadFont(char, font)
			if err != nil {
				return nil
			}

			chars = append(chars, ascii)
		}

		fmt.Println(len(chars), len(word))
		output = append(output, chars)
	}

	return output
}
