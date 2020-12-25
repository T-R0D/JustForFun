package input

import (
	"io/ioutil"
	"strings"
)

func ReadInput(inputPath string) (string, error) {
	bytes, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return "", err
	}

	input := string(bytes)

	return strings.TrimSpace(input), nil
}
