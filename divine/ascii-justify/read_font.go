package asciijustify

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (j *Justify) ReadFont(char rune, font string) ([]string, error) {
	path := ""

	switch strings.ToLower(font) {
	case "shadow":
		path = "assets/shadow.txt"
	case "standard":
		path = "assets/standard.txt"
	case "tinkertoy":
		path = "assets/tinkertoy.txt"
	default:
		path = "assets/standard.txt"
	}

	startLine := ((char - ' ') * 9) + 1
	endLine := startLine + 8
	currentLine := 0

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	var output []string

	for scanner.Scan() {
		if currentLine >= int(startLine) && currentLine <= int(endLine) {
			output = append(output, scanner.Text())
		} else if currentLine > int(endLine) {
			break
		}
		currentLine += 1
	}

	if err = scanner.Err(); err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return output, nil
}
