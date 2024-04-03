package main

import (
	"fmt"
	"time"
)

func main() {

	SendSignForChannel()
}

func SendSignForChannel() {
	mutexChannel := make(chan bool)
	otherChannel := make(chan bool, 200)
	go tryGetMutexChannelSign(mutexChannel)
	for i := 0; i < 100; i++ {
		trySendBlockSign(mutexChannel, otherChannel)
	}
	time.Sleep(8 * time.Second)
	trySendBlockSign(mutexChannel, otherChannel)
	time.Sleep(60 * time.Second)
}

func tryGetMutexChannelSign(mutexChannel chan bool) {
	<-mutexChannel
}

// 处理任务
func doTask(mutexChannel chan bool) {
	// 释放blockChannel中的消息
	defer tryGetMutexChannelSign(mutexChannel)
	fmt.Println("正在处理任务")
	time.Sleep(5 * time.Second)
	fmt.Println("任务处理完成")
}

func trySendBlockSign(mutexChannel chan bool, otherChannel chan bool) {
	select {
	case mutexChannel <- true:
		fmt.Println("我发送了sign给blockChannel")
		go doTask(mutexChannel)
	case otherChannel <- true:
		fmt.Println("sign去其他Channel了")
	}
}
