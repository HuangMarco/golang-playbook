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

- 记住：默认情况下声明的channel都是unbuffered的
- unbuffered channel表示每一次single send(ch <-)将会被block，直到有goroutine执行receive动作(<- ch)
- 如果buffered channel的capacity为1，表明

By default, a channel has a buffer size of 0 (you get this with make(chan int)). This means that every single send will block until another goroutine receives from the channel. A channel of buffer size 1 can hold 1 element until sending blocks, so you'd get

```go
c := make(chan int, 1)
c <- 1 // doesn't block
c <- 2 // blocks until another goroutine receives from the channel
```

关于上面也可以查看：https://golang.org/doc/effective_go#channels

同时一个很好的例子说明：

```go
package main

import (
    "fmt"
    "time"
)

func receiver(ch <-chan int) {
    time.Sleep(500 * time.Millisecond)
    msg := <-ch
    fmt.Printf("receive messages  %d from the channel\n", msg)
}

func main() {
    start := time.Now()
    zero_buffer_ch := make(chan int, 0)
    go receiver(zero_buffer_ch)
    zero_buffer_ch <- 444
    elapsed := time.Since(start)    
    fmt.Printf("Elapsed using zero_buffer channel: %v\n", elapsed)

    restart := time.Now()
    non_zero_buffer_ch := make(chan int, 1)
    go receiver(non_zero_buffer_ch)
    non_zero_buffer_ch <- 4444
    reelapsed := time.Since(restart)
    fmt.Printf("Elapsed using non zero_buffer channel: %v\n", reelapsed)
}
```


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

- 通过select关键字，可以对goroutine(也就是channel)做swith...case操作
- 与switch case不同的是，select会一直等待直到所有的channel都ready
- select通常与time package配合，使用场景：通常当你需要同步不同的操作的时候，使用select等待所有channel同时完成，从而继续到下一步

```golang
package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

同样查看test-channel-select.go

### select - default

- 当select block中没有任何一个case ready的时候，执行default

```golang
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

### exercise - Equivalent Binary Trees

https://gist.github.com/kaipakartik/8120855

## golang - sync

https://tour.golang.org/concurrency/9

如果说之前的select等都是为了让多个goroutine同时执行，那么sync的作用就是：一次只能允许一个协程运行，或者一次只能允许一个goroutine访问某个变量

这个概念被称为mutual exclusion，为了简便来说，称之为mutex.记住哈，这只是个概念，而不是一个类库。

golang针对于该概念提供类库:[sync.Mutex](https://golang.org/pkg/sync/#Mutex)且支持以下两种功能：

- lock
- unlock

A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex.
Mutex是一个共有排他锁，0值代表unlocked状态

A Mutex must not be copied after first use.

```golang
func (m *Mutex) Lock()
// goroutine调用该方法来给m上锁，如果锁已经被使用了，那么该goroutine将会被阻塞直到mutex可以使用。一个被上了锁的mutex不属于任何一个goroutine，goroutine只能执行上锁解锁动作，goroutine a给mutext上锁，然后golang安排另外一个goroutine给该mutex解锁

func (m *Mutex) Unlock()
// goroutine调用该方法给m解锁，如果m本身没有被锁住，就会报runtime error
```


