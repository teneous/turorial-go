package exception

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

//网络异常，编译声明异常
func ioErr() (string, error) {
	_, err := ioutil.ReadDir("ioexception.go")
	return "empty", err
}

//运行时异常，手动判断？
func algErr() (int, error) {
	rand.Seed(time.Now().UnixNano())
	total := rand.Intn(1)
	sum := 1

	if total == 0 {
		err := fmt.Errorf("1/0 %s: unsupported operation exception)", "ioexception.go")
		return -1, err
	} else {
		return sum / total, nil
	}
}
