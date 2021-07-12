package basic

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

//test read file
func TestReadFile(t *testing.T) {
	file, err := os.Open("../../resources/file.properties")
	if err != nil {
		fmt.Printf("read file encounter exception")
		return
	} else {
		reader := bufio.NewReader(file)
		println("start reading file properties")
		for true {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			} else {
				fmt.Printf("reading property %s\n", string(line))
			}
		}
	}
}
