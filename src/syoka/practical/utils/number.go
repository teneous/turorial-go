package utils

//IsNotZero is check parameter is not equals to zero
func IsNotZero(i int) bool {
	return i != 0
}

func IsPositive(i int) bool {
	return i >= 0
}

func IsNegative(i int) bool {
	return !IsPositive(i)
}
