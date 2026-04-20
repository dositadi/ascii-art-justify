package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lineWidth = 0
var termWidth = os.Getenv("COLUMNS")
var termWidthInt, _ = strconv.Atoi(termWidth)

func ProcessInput(option string, strVal string, banner string) (final string, err error) {

	var rawLetters []string
	path := "./" + "assets" + "/" + banner + ".txt"
	file, err := os.Open(path)

	if err != nil {
		return "", fmt.Errorf("Invalid 1")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rawLetters = append(rawLetters, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return "", fmt.Errorf("Invalid 2")
	}

	processedLetters := chunkSlice(rawLetters, 9)
	strValue := strings.Split(strVal, "\\n")

	for i, strVal := range strValue {
		var holdWord [][]string

		for _, str := range strVal {
			index := int(str) - 32
			holdWord = append(holdWord, processedLetters[index])
		}
		final += printChar(holdWord, option)
		if i == 0 {
			lineWidth = getCharWidth(holdWord)
		}
	}

	return final, nil
}

func chunkSlice(sliceVal []string, chunkSize int) (hold [][]string) {
	if chunkSize <= 0 {
		return [][]string{}
	}

	for i := 0; i < len(sliceVal); i += chunkSize {
		end := len(sliceVal)
		if i+chunkSize <= len(sliceVal) {
			end = i + chunkSize
		}
		hold = append(hold, sliceVal[i:end])
	}

	return
}

func getCharWidth(hold [][]string) int {
	var buildString string
	for i := 0; i < len(hold); i++ {
		buildString += hold[i][0]
	}
	return len(buildString)
}

func printChar(hold [][]string, option string) (buildString string) {
	//assume the length of i is 9
	var temp string
	for i := 0; i < 9; i++ {
		for j := 0; j < len(hold); j++ {
			temp += hold[j][i]
		}
		if option == "right" {
			diff := termWidthInt - lineWidth
			buildString += strings.Repeat(" ", diff) + temp + "\n"
			temp = ""
		}
	}
	return
}
