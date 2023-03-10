package util

import "strings"

func TrimNewLine(inputStr string) string {
	return strings.TrimRight(string(inputStr), "\n")
}
