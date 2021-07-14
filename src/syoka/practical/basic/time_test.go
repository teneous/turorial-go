//用于测试时间类
//15 hour,04 min 05 sec
package basic

import (
	"fmt"
	"testing"
	"time"
)

const DF = "2006-01-02 15:04:05"

func TestTimeParse(t *testing.T) {
	parseTime, err := time.Parse(DF, "2020-01-01 12:00:30")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println(parseTime)
	}

}
