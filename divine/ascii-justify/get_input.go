package asciijustify

import (
	"os"
	"strings"
)

func (j *Justify) GetUserInput() (string, string, string) {
	if len(os.Args) <= 2 {
		j.PrintError(UsageErr)
		return "", "", ""
	}

	args := os.Args[1:]

	if !strings.HasPrefix(args[0], AlignFlag) {
		j.PrintError(UsageErr)
		return "", "", ""
	}

	alignValue := strings.TrimPrefix(args[0], AlignFlag)

	if !((alignValue == "right") || (alignValue == "left") || (alignValue == "center") || (alignValue == "justify")) {
		j.PrintError(UsageErr)
		return "", "", ""
	}

	var alignment, text, font string

	switch len(args) {
	case 2:
		alignment = alignValue
		text = args[1]
		font = "standard"
	case 3:
		alignment = alignValue
		text = args[1]
		font = args[2]
	default:
		j.PrintError(UsageErr)
		return "", "", ""
	}
	return alignment, text, font
}
