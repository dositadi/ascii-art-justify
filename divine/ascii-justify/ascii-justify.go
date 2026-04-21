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
	_, _, err := j.GetTerminalSize()
	if err != nil {
		j.PrintError(TerminalSizeError + " " + err.Error())
		return
	}

	_, _, _ = j.GetUserInput()
}

func (j *Justify) PrintError(err string) {
	fmt.Println(err)
}
