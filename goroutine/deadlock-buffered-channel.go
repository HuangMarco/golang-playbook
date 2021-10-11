package main

import (
	"fmt"
	// "time"
)

//deadlock：程序因为运行期panic导致错误的情况
func main(){
	ch := make(chan string, 2)

	ch <- "test1"
	ch <- "test2"
	//此处将会报错，因为只有往channel里面塞，却没有读取操作
	ch <- "test3"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan send]:
// main.main()
// 	/Users/i323691/work_dir/training/go-training/golang-playbook/goroutine/deadlock-buffered-channel.go:14 +0x8d
// exit status 2
// i323691@C02Z23J2LVDR goroutine %