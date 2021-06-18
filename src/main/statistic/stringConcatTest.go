package main

import (
	"bytes"
	"fmt"
	"time"
)

//测试bytebuffer和str的字符串拼接
//测试结果，1000次 大约25倍
//测试结果，10000次 大约100倍
//测试结果，100000次 大约750倍
func main() {
	normalLoop("z")
	println()
	bufferLoop("c")
}

func normalLoop(dupStr string) {
	start := time.Now()
	var result = ""
	for i := 0; i < 100000; i++ {
		result += dupStr
	}
	end := time.Now()

	duration := end.UnixNano() - start.UnixNano()
	fmt.Printf("normalLoop: 执行总时间:%d,结果长度%d", duration, len(result))
}

func bufferLoop(dupStr string) {
	start := time.Now()
	buffer := bytes.Buffer{}
	var result = ""
	for i := 0; i < 100000; i++ {
		buffer.WriteString(dupStr)
	}
	result = buffer.String()
	end := time.Now()

	duration := end.UnixNano() - start.UnixNano()
	fmt.Printf("bufferLoop: 执行总时间:%d,结果长度%d", duration, len(result))
}
