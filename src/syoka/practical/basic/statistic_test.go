//测试bytebuffer和str的字符串拼接
//测试结果，1000次 大约25倍
//测试结果，10000次 大约100倍
//测试结果，100000次 大约750倍
package basic

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestNormalLoop(t *testing.T) {
	var str string = "z"
	start := time.Now()
	var result = ""
	for i := 0; i < 100000; i++ {
		result += str
	}
	end := time.Now()

	duration := end.UnixNano() - start.UnixNano()
	fmt.Printf("normalLoop: 执行总时间:%d,结果长度%d", duration, len(result))
}

func TestBufferLoop(t *testing.T) {
	var str string = "z"
	start := time.Now()
	buffer := bytes.Buffer{}
	var result = ""
	for i := 0; i < 100000; i++ {
		buffer.WriteString(str)
	}
	result = buffer.String()
	end := time.Now()

	duration := end.UnixNano() - start.UnixNano()
	fmt.Printf("bufferLoop: 执行总时间:%d,结果长度%d", duration, len(result))
}
