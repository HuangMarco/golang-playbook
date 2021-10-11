# Goroutine

- goroutine相当于收发消息的通道
- 其方向可以双向可以单向
- 类似于消息机制，其中的消息可以被buffered

## sync

Go 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发。

例如我们希望并发下载 N 个资源，多个并发协程之间不需要通信，那么就可以使用 sync.WaitGroup，等待所有并发协程执行结束。

## golang goroutine

```golang
package main

import "fmt"

func f(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
  }
}

func main() {
  go f(0)
  var input string
  fmt.Scanln(&input)
}
```

在上面的代码中：

* 有2个routine: main函数本身；go f(0)
* 当执行某个函数的时候，函数中所有代码都会按顺序逐行执行，如果某行是调用某个函数，那么会进入到该函数中，执行所有该函数中的代码，当该行函数调用执行完成之后立刻进入到下一行。
* 如果是goroutine,那么不会进入到某行对应的函数中，并不会等待该函数执行完毕，而是立刻进入到下一行，参照上面用例，不会等待f(0)执行完成，而是立刻进入到var input string
* goroutine是轻量级的，所以可以同时起多个goroutine,如下所示

```golang
func main() {
  for i := 0; i < 10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
```

### add delay

```golang
package main

import (
  "fmt"
  "time"
  "math/rand"
)

func f(n int) {
  for i := 0; i < 10; i++ {
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(250))
    time.Sleep(time.Millisecond * amt)
  }
}

func main() {
  for i := 0; i < 10; i++ {
    go f(i)
  }
  var input string
  fmt.Scanln(&input)
}
```

## golang channel

https://tour.golang.org/concurrency/2

```golang
ch <- v    // Send v to channel ch. 发送方
v := <-ch  // Receive from ch, and 接收方
           // assign value to v.
```

```golang
// 先声明再使用
ch := make(chan int)
```

默认情况下，sends和recieves block（对应相应的代码行）都会处于block状态直到另外一方完全ready才正式启用。有了这个机制，就可以使得goroutines之间可以达到同步的效果，也不需要像java那样使用lock或者conditional values这种

https://www.golang-book.com/books/intro/10


channel的主要作用：

- 使2个goroutine能够彼此通信协调
- 同步goroutine之间的执行

下面的例子：

```golang
package main

import (
  "fmt"
  "time"
)

func pinger(c chan string) {
  for i := 0; ; i++ {
    c <- "ping"
  }
}

func printer(c chan string) {
  for {
    msg := <- c
    fmt.Println(msg)
    // 上面2行可以简写为：fmt.Println(<-c)
    time.Sleep(time.Second * 1)
  }
}

func ponger(c chan string) {
  for i := 0; ; i++ {
    c <- "pong"
  }
}

func main() {
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
```

从上述代码中：

- channel的声明是通过关键字chan来完成
- 如果要传递什么样类型的消息就使用该类型作为声明类型，比如上述的make(chan string)就是声明了一个字符串类型的channel，传递给channel的是字符串类型
- `<-`表示send and receive message on the channel
- `c <- "ping"`表示将字符串ping发送到channel中
- `msg := <- c`表示将接收到的channel中的字符串内容存储在msg变量中
- `fmt.Println(msg)`可以改写为`fmt.Println(<-c)`这样就可以将msg声明语句与fmt打印语句合并为一条语句
- `go pinger(c)`和`go printer(c)`通过同一个channel c来达到同步2个routine的目的，当pinger希望发送消息的时候，它会等待printer准备好了接收消息之后，才会发送消息，这个相当于blocking

### Channel direction 限定channel的方向

我们可以限定channel的方向，从而限定其要么只能seding，要么只能receiving message

```golang
// 限定pinger只能作为发送方发送消息
func pinger(c chan <- string)
```

### buffered channel

```golang
ch := make(chan int, 100)
```

详情见同目录`buffered-channel.go`

### Deadlocks in Buffered Channel

`deadlock-buffered-channel.go`

### range and close

- sender可以关闭一个channel，将会导致不再有值会塞入到channel中
- 永远只能由sender关闭channel，receiver无权关闭channel
- receiver可以测试channel
- channel被关闭之后，再次尝试向channel中塞入值，会导致panic
- 一般来说channel无需手动关闭，除非你需要显式告诉receiver channel中已经不会有新的值了

```go
ch := make(chan int)
//可以通过下面语句嗅探通道是否关闭
v, ok := <- ch

//ok为false的时候，说明channel中没有新的值需要读取了，或者channel被关闭
```

golang channel range操作：

```golang
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

```

## golang channel select

- goroutine可以通过