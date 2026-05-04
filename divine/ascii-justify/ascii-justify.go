package asciijustify

import (
	"fmt"
)

const (
	UsageErr          = "Usage: go run . [OPTION] [STRING] [BANNER]\nExample: go run . --align=right something standard"
	TerminalSizeError = "Cant get terminal size!"
	AlignFlag         = "--align="
)

type Justify struct{}

func (j *Justify) Run() {
	_, text, font := j.GetUserInput()

	_, width, err := j.GetTerminalSize()
	if err != nil {
		j.PrintError(TerminalSizeError + " " + err.Error())
		return
	}

	splitted := j.SplitByNewLine(text)

	transformed := j.TransformAndJustify(splitted, font)

	j.PrintAscii(transformed, width)
}

func (j *Justify) PrintError(err string) {
	fmt.Println(err)
}
