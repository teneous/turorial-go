package utils

func ArrayLen(source []interface{}) int {
	if source == nil {
		return -1
	} else {
		return len(source)
	}
}

func ArrayStringLen(source []string) int {
	if source == nil {
		return -1
	} else {
		return len(source)
	}
}
