package asciijustify

import (
	"fmt"
	"strings"
)

const (
	UsageErr          = "Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard"
	TerminalSizeError = "Cant get terminal size!"
	AlignFlag         = "--align="
)

type Justify struct{}

func (j *Justify) Run() {

	_, text, font := j.GetUserInput()

	if !strings.Contains(text, " ") {
		splitted := j.SplitByNewLine(text)

		transformed := j.NormalTransform(splitted, font)

		j.PrintAscii(transformed, 0)
	} else {
		_, width, err := j.GetTerminalSize()
		if err != nil {
			j.PrintError(TerminalSizeError + " " + err.Error())
			return
		}

		fmt.Println(width)

		splitted := j.SplitByNewLine(text)

		transformed := j.TransformAndJustify(splitted, font)

		j.PrintAscii(transformed, width)
	}

	fmt.Println("spaces: ", j.calcAmountOfSpace(166, [][]string{{"Abeg", "go", "school", "there"}, {"drop"}, {"am"}}))
}

func (j *Justify) PrintError(err string) {
	fmt.Println(err)
}
