package main

import (
	"fmt"
	// "time"
)

// 这样的声明限定函数pinger只能发送消息
func pinger(c chan <- string){
	for i:= 0;i<=10 ;i++ {
		// 发送ping
		c <- "ping"
	}
}

// 这样的声明限定函数ponger只能接受消息
func ponger(c <-chan string) {
	for i:=0;i<=10 ;i++ {
		// 发送ping
		msg := <-c
		fmt.Println("received from pinger: ",msg)
	}
}

// func printer(c chan string){
// 	for i:=0; ;i++ {
// 		// 接收方接收消息
// 		fmt.Println(<-c)
// 		time.Sleep(time.Second * 1)
// 	}
// }

func main(){
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	// go printer(c)

	var input string
	fmt.Scanln(&input)

	//尝试更改capacity，由0更改到1，更改到2查看效果
	ch := make(chan int,2)
	ch <- 1
	ch <- 2
}