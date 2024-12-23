package Tools

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Lines(s string) []string {
	return strings.Split(s, "\n")
}

func ReadByLines(path string) []string {
	file, err := os.Open(path)

	// We don't expect wrong file paths when every input is a static file
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	lines := make([]string, 0)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func Read(path string) string {
	file, err := os.ReadFile(path)

	// We don't expect wrong file paths when every input is a static file
	if err != nil {
		fmt.Println(err)
	}
	input := string(file)

	return strings.Trim(input, " \n")
}
