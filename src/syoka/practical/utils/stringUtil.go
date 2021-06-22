package utils

import "strings"

func IsEmpty(str string) bool {
	return len(strings.ReplaceAll(str, " ", "")) != 0
}

func IsNotEmpty(str string) bool {
	return !IsEmpty(str)
}
