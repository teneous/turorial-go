//生产者消费者模型 channel中的数据只会被一个消费者消费
package concurrency

import (
	"fmt"
	"testing"
	"time"
)

//TestConsumerAndProvider 测试生产者,消费者通过channel传递消息
func TestConsumerAndProvider(t *testing.T) {
	//create channel
	channel := make(chan string)

	go producer(channel)
	go consumer("zhangsan", channel)
	go consumer("lisi", channel)
	time.Sleep(3 * time.Second)
}

//数据上有
func producer(channel chan string) {
	fmt.Println("钢棍谢师傅开始制作")
	channel <- "烧卖 1份"
	fmt.Println("钢棍谢师傅完成制作")
}

func consumer(name string, channel chan string) {
	fmt.Printf("消费者 %s 开始消费\n", name)
	foods := <-channel
	fmt.Printf("消费者 %s 品尝了%s\n", name, foods)
}
