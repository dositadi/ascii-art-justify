package asciijustify

func (j *Justify) TransformAndJustify(input []string, font string) [][][]string {
	var output [][][]string

	for _, word := range input {
		var chars [][]string

		if word == "" {
			chars = append(chars, []string{""})
			output = append(output, chars)
			continue
		}

		for _, char := range word {
			if char < ' ' || char > '~' {
				continue
			}

			if char == ' '{
				space := []string{" "}
				chars = append(chars, space)
				continue
			}

			ascii, err := j.ReadFont(char, font)
			if err != nil {
				return nil
			}

			chars = append(chars, ascii)
		}

		output = append(output, chars)
	}

	return output
}
