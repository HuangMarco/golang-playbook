package main

import(
	"fmt"
	"time"
)

func write(ch chan int){
	for i:=0 ; i <=4 ; i++{
		//依次将i写入到ch中
		//因为容量限定为2，所以只能立即将0和1立即写入到ch中，然后阻塞等待，直到0或1中的任何一个被读取之后，才能继续写入值到ch中
		ch <- i
		fmt.Println("成功将 %T写入到ch", i)
	}
	// 关闭ch
	close(ch)
}

func main(){
	// 创建一个容量为2的channel
	
	ch := make(chan int, 2)

	go write(ch)
	//等待10秒 
	time.Sleep(10*time.Second)
	// channel也是可以被循环遍历的
	for v:= range ch{
			fmt.Println("收到value %T", v, "from ch")
			//每次读完1个之后，停留4秒读取下一个
			time.Sleep(4*time.Second)
	}
}

// 打印结果
// 成功将 %T写入到ch 0
// 成功将 %T写入到ch 1
//此处停留10秒，只有一次停留10秒
// 收到value %T 0 from ch
//停留4秒
// 成功将 %T写入到ch 2
// 收到value %T 1 from ch
//停留4秒
// 成功将 %T写入到ch 3
// 收到value %T 2 from ch
//停留4秒
// 成功将 %T写入到ch 4
// 收到value %T 3 from ch
// 收到value %T 4 from ch

//可见：容量就是同一时间在channel中存在的消息的数量
