package main

import "fmt"

/**
 * zhi
 */
func main() {
	cmd := "gh clone repo"
	showMemoryAddress(cmd)
	showMemoryAddress2(&cmd)
	showMemoryAddress3(&cmd)
	showMemoryValue(&cmd)
}

func showMemoryAddress(cmd string) {
	fmt.Println(&cmd)
}

func showMemoryAddress2(cmd *string) {
	fmt.Println(cmd)
}

func showMemoryAddress3(cmd *string) {
	fmt.Println(&cmd)
}

func showMemoryValue(cmd *string) {
	fmt.Println(*cmd)
}
