package main

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {

	result := map[rune][]string{}

	if !strings.HasSuffix(filename, ".txt") {
		filename += ".txt"
	}

	data, _ := os.ReadFile("banners/" + filename)

	lines := strings.Split(strings.ReplaceAll(string(data), "\\n", "\n"), "\n")

	start := ' '
	stop := '~'

	for ch := start; ch <= stop; ch++ {
		startIndex := int(ch-start)*9 + 1
		result[ch] = lines[startIndex : startIndex+8]
	}

	return result, nil
}
