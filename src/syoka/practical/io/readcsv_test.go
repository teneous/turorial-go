package io

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestReadCsv(t *testing.T) {

	var path = "../../resources/student.csv"
	_, err := os.Stat(path)
	if err != nil {
		println("no file found.err:", err)
	} else {
		f, err := os.Open(path)
		defer f.Close()
		if err != nil {
			println("can't open file.err:", err)
		} else {
			reader := csv.NewReader(f)
			records, err := reader.ReadAll()
			if err != nil {
				println("csv read record failed.err:", err)
			} else {
				for _, val := range records {
					fmt.Println(val)
				}
			}
			reader.FieldsPerRecord = -1

		}
	}
}

func TestWriteCsv(t *testing.T) {
	err := ioutil.WriteFile("demo.csv", []byte("name,age"), 0777)
	if err != nil {
		println("creat file failed")
	}
}
