package asciijustify

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func (j *Justify) GetTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")

	cmd.Stdin = os.Stdout

	result, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}

	var width, height int

	sizes := strings.Fields(string(result))

	height, _ = strconv.Atoi(sizes[0])
	width, _ = strconv.Atoi(sizes[1])

	return height, width, nil
}
