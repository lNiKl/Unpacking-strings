package main

import (
	"fmt"
	"strconv"
)

func unpacking(word string) string {
	unpacked := ""
	if isDigit(rune(word[0])) {
		return "'' (некорректная строка)"
	}
	for i := 0; i < len(word); i++ {
		switch {
		case word[i] == '\\':
			if i == len(word)-1 {
				return unpacked
			} else {
				unpacked += string(word[i+1])
				i++
			}
		case isDigit(rune(word[i])):
			k, _ := strconv.Atoi(string(word[i]))
			for j := 0; j < k-1; j++ {
				unpacked += string(word[i-1])
			}
		default:
			unpacked += string(word[i])
		}
	}
	return unpacked
}
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func main() {
	inputs := []string{"a4bc2d5e", "abcd", "45", "qwe\\4\\5", "qwe\\45", "qwe\\5", "qwe\\"}

	for _, input := range inputs {
		fmt.Println(input + " => " + unpacking(input))
	}
}
