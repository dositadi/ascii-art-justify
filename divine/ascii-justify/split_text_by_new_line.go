package asciijustify

import "strings"

func (j *Justify) SplitByNewLine(text string) []string {
	return strings.Split(text, "\\n")
}
