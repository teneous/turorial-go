package main

import "strings"

func isEmpty(str string) bool {
	return len(strings.ReplaceAll(str, " ", "")) != 0
}

func isNotEmpty(str string) bool {
	return !isEmpty(str)
}
