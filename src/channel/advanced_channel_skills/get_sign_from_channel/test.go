package main

import (
	"fmt"
	"time"
)

func main() {
	getSignFromChannel()
}

func getSignFromChannel() {
	mutexChannel := make(chan bool)
	otherChannel := make(chan bool, 200)
	// 在程序启动的时候就发送消息
	go sendMutexChannelSign(mutexChannel)

	for i := 0; i < 100; i++ {
		tryGetMutexSign(mutexChannel, otherChannel)
	}
	time.Sleep(5 * time.Second)
	// 尝试又发送消息给blockChannel
	go sendMutexChannelSign(mutexChannel)
	fmt.Println("-------")
	tryGetMutexSign(mutexChannel, otherChannel)
	time.Sleep(60 * time.Second)
}

func sendMutexChannelSign(mutexChannel chan bool) {
	mutexChannel <- true
}

// 处理任务
func doTask(mutexChannel chan bool) {
	// 发送 mutexChannel 中的消息
	defer sendMutexChannelSign(mutexChannel)
	fmt.Println("正在处理任务")
	time.Sleep(5 * time.Second)
	fmt.Println("任务处理完成")
}

func tryGetMutexSign(mutexChannel chan bool, otherChannel chan bool) {
	select {
	case <-mutexChannel:
		fmt.Println("我拿到了blockChannel中的sign")
		// 处理任务
		go doTask(mutexChannel)
	default: // mutexChannel 被阻塞
		otherChannel <- true
		fmt.Println("sign去其他Channel了")
	}
}
