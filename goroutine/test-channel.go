package main

import "fmt"

func pinger(c chan int){

	for i:=0;i<=10 ;i++{
		// fmt.Println("pinger",i)
		c <- int(i)
	}

}

func ponger(c chan int){
	for i:=0;i<=10 ;i++{
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {

	ch := make(chan int, 1)
	
	go pinger(ch)
	go ponger(ch)
	
	//下面这段代码如果被注释掉，那么从命令行上看起来，代码执行起来不会有任何打印信息，料想到的fmt.Println(msg)不会在命令后打印出来
	//从routine角度看起来，很正常，因为go pinger不会等待pinger执行结束才会到下一行，而是立刻到下一行执行
	// fmt.Println("gtset")
	// fmt.Println(<-ch)
	var input string
	fmt.Scanln(&input)
}
