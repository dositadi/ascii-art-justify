package utils

import (
	"fmt"
	"slices"
)

var validOptions, validBanners []string = []string{"left", "right", "center", "justify"}, []string{"shadow", "tinkertoy", "standard"}

func ReceiveInput(option string, strVal string, banner string) (string, error) {

	//Processing options and banners
	if option == "" {
		option = "right"
	}

	if banner == "" {
		banner = "standard"
	}

	if !slices.Contains(validOptions, option) || !slices.Contains(validBanners, banner) {
		return "", fmt.Errorf("Invalid Input")
	}
	str, err := ProcessInput(option, strVal, banner)
	if err != nil {
		return "", err
	}
	//

	return str, nil

}
