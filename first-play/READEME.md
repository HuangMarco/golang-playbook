# 第一轮-Round One - 最初级菜鸟阶段

https://github.com/goinaction/code

第一件事：记住golang的官方网站：https://golang.org

## 安装

```sh
# linux ubuntu环境：https://golang.org/doc/install?download=go1.16.4.linux-amd64.tar.gz

wget https://golang.org/dl/go1.16.4.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

## 概念

最大限度利用别人已有的网站资源，归为己用

## golang func

一个典型的函数定义如下，使用关键字 func，参数可以有多个，返回值也支持有多个。特别地，package main 中的func main() 约定为可执行程序的入口。

```golang
func funcName(param1 Type1, param2 Type2, ...) (return1 Type3, ...) {
    // body
}


// 例如，实现2个数的加法（一个返回值）和除法（多个返回值）：
func add(num1 int, num2 int) int {
	return num1 + num2
}

func div(num1 int, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}
func main() {
	quo, rem := div(100, 17)
	fmt.Println(quo, rem)     // 5 15
	fmt.Println(add(100, 17)) // 117
}
```

**也可以给返回值命名，简化 return，例如 add 函数可以改写为**

```golang
func add(num1 int, num2 int) (ans int) {
	ans = num1 + num2
	return
}
```

##  golang method

- golang method是一种特殊类型的golang function
- 通常需要面临的情况，就是不同的函数都需要处理同一份data，每次处理逻辑不同，需要使用的参数也不同
- golang通过methods来定义特殊类型的func，目的是为了处理golang中的一些特殊的数据类型：receiver

https://www.digitalocean.com/community/tutorials/defining-methods-in-go

```golang
package main

import "fmt"

//创建了struct类型的变量：Creature
type Creature struct{
    //Creature类型有2个字段
    Name string
    Greeting string
}

//声明和普通函数类似
//只是在func关键字后面加上特殊的参数：指定method的receiver
//receiver概念类似于java的class，即某个receiver拥有某个方法。以下面为例：
//Creature类型拥有Greet方法
//至此，Creature有了方法Greet
//社区的命名规则：首字母变小写，即creature,但是你可以随意命名他，但是还是最好按照规范来
func (c Creature) Greet(){
    fmt.Printf("%s says %s", c.Name, c.Greeting)
}

func main(){
    //创建Creature类型的变量marco
    marco := Creature{
        Name: "Marco",
        Greeting: "Welcome to the golang world!",
    }
    //第一种调用Greet方法的方式
    Creature.Greet(marco)
    //第二种调用方式，与第一种效果相同
    //社区推崇这种调用方式：通过dot notation来调用方法，使用到变量marco中默认golang为我们存储的Creature
    marco.Greet()
}

//output:
//Marco says Welcome to the golang world!

```

```golang
//第二种方式证明为什么建议使用第二种方式来做golang中的方法调用

package main

import "fmt"

type Creature struct{
    Name string
    Greeting string
}

//为类型Creature定义方法Greet同时返回一个Creature
func (creature Creature) Greet() Creature{
    fmt.Printf("%s says %s !\n", c.Name, c.Greeting)
}
//为类型Creature定义方法SayGoodbye，接收一个参数
func (creature Creature) SayGoodbye(name string){
    fmt.Println("Farewell", name, "!")
}

func main(){
    marco := Creature{
        Name: "Marco",
        Greeting: "Hello!"
    }

    //这种调用方式更加高效，但是同时要求方法Greet返回一个Creature
    marco.Greet().SayGoodbye("golang")
    //比较低效的证明
    Creature.SayGoodbye(Creature.Greet(marco), "golang")
}

//output:
//Marco says Hello!
//Farewell golang!
//Marco says Hello!
//Farewell golang!

```

### Pointer/Value Receivers 声明- methods on values/pointers

- 分为methods on values和methods on pointers: https://golang.org/doc/faq#methods_on_values_or_pointers
- 考虑的重点是：你是否有需求去修改传人类型的本身

```golang

func (s *MyStruct) pointerMethod() { } // method on pointer
func (s MyStruct)  valueMethod()   { } // method on value

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    //...
}
```

```golang

package main
 
import "fmt"
 
type Mutatable struct {
    a int
    b int
}

//不会改变receiver
func (m Mutatable) StayTheSame() {
    m.a = 5
    m.b = 7
}
//会改变receiver
func (m *Mutatable) Mutate() {
    m.a = 5
    m.b = 7
}
 
func main() {
    m := &Mutatable{0, 0}
    fmt.Println(m)
    m.StayTheSame()
    fmt.Println(m)
    m.Mutate()


//output:
//&{0 0}
//&{0 0}
//&{5 7}
```

一般来说，指针通常在函数传递参数，或者给某个类型定义新的方法时使用。Go 语言中，参数是按值传递的，如果不使用指针，函数内部将会拷贝一份参数的副本，对参数的修改并不会影响到外部变量的值。如果参数使用指针，对参数的传递将会影响到外部变量。

例如：

```golang
func add(num int) {
	num += 1
}

func realAdd(num *int) {
	*num += 1
}

func main() {
	num := 100
	add(num)
	fmt.Println(num)  // 100，num 没有变化

	realAdd(&num)
	fmt.Println(num)  // 101，指针传递，num 被修改
}
```

## golang引用类型-interface

- golang中，interface类型和java中的接口功能类似，如果声明了interface类型，在interface类型中声明所有需要被实现的方法，带调用时，当某种类型的变量被指定为interface类型，那么该类型必须实现interface中声明的所有方法
- 一旦将某个类型指定为interface类型，那么golang就会检查该类型所有的方法集合，查看是否实现interface里面的方法
- golang interface声明

案例1:

```golang
package main

import "fmt"

type Person interface{
	getName() string
}

type Student struct{
	age int
	name string
}

//这里不会改变stu.name，所以声明指针类型方法pointer receiver还是value receiver都无所谓
func (stu *Student) getName() string{
	return stu.name
}

func main(){
	xiaoming := &Student{
		age : 18,
		name: "xiaoming",
	}
	
	fmt.Println(xiaoming.getName())
}

// print: xiaoming
```

案例2:
```golang

package main

import {

    "fmt"
    "strings"
}

type Stringer interface {
    //为该接口Stringer声明方法String()
    String() string
}

//声明一个结构类型，名为Ocean
type Ocean struct{
    Creatures []string
}

//结构类型Ocean声明并实现方法String()
func (ocean Ocean) String() string {
    return strings.Join(ocean.Creatures, ", ")
}

func log(header string, s fmt.Stringer){
    fmt.Println(header, ":", s)
}

func main(){
    ocean := Ocean{
        Creatures: []string{
            "sea urchin",
            "lobster",
            "shark",
        }
    }
}

log("The creatures in ocean:", ocean)

//如果Ocean没有实现String()方法，就会报错
// src/e4/main.go:24:6: cannot use o (type Ocean) as type fmt.Stringer in argument to log:
        // Ocean does not implement fmt.Stringer (missing String method)
```

### 检测structure是否实现interface定义的方法

如何确保某个类型实现了某个接口的所有方法呢？一般可以使用下面的方法进行检测，如果实现不完整，编译期将会报错。

```golang
var _ Person = (*Student)(nil)
```

* 将空值 nil 转换为 *Student 类型，再转换为 Person 接口，如果转换失败，说明 Student 并没有实现 Person 接口的所有方法。

### 实例与接口之间的转换

实例可以强制类型转换为接口，接口也可以强制类型转换为实例。

```golang
func main() {
	var p Person = &Student{
		name: "Tom",
		age:  18,
	}

	stu := p.(*Student) // 接口转为实例
	fmt.Println(stu.getAge())
}

```

- pointer receiver的methods set与value receiver的methods set是不同的，因为pointer receiver可以改动receiver，但是value receiver不能。只有pointer receiver才能满足interface的需要

```golang
package main

import "fmt"

type Submersible interface {
    Dive()
}

type Shark struct {
    Name string
    isUnderwater bool
}

//定义了value receiver
func (s Shark) String() string {
    if s.isUnderwater {
        return fmt.Springtf("%s is underwater", s.Name)
    }
    return fmt.Springtf("%s is on the Surface", s.Name)
}

//结构体Shark因为要调用接口Submersible中的方法Dive，所以必须实现方法Dive
//同时是pointer reciever，所以改变了结构体自身
func (s *Shark) Dive() {
    s.isUnderwater = true
}

//方法submerge中，参数为接口类型，所以传入的参数必须实现接口Submersible中的所有方法
func submerge(s Submersible) {
    //接口类型又通过调用方法Dive，改变了receiver自身
    s.Dive()
}

func main() {
    //因为当前情况下，Shark类型实现的point receiver方式的接口Dive，所以在这里也必须传入一个指针类型变量
    //如果传入的不是一个指针类型变量，就会报错
    s := &Shark{
        Name: "Marco",
    }

    fmt.Println(s)
    submerge(s)
    fmt.Println(s)
}

// Output
// Marco is on the surface
// Marco is underwater

```

### golang标识符

https://haicoder.net/golang/golang-identifier.html

- Golang 对各种 变量、方法、函数 等命名时使用的字符序列称为标识符。它的主要作用就是作为变量、函数、类、模块以及其他对象的名称
- 凡是自己可以起名字的地方都叫标识符
- Go 语言包名需要和目录保持一致，尽量采取有意义的包名，简短，有意义，不要和标准库包名冲突
- Go 语言中的变量名、函数名、常量名 都需要采用驼峰法
- 如果变量名、函数名、常量名首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用( 注：可以简单的理解成，首字母大写是公开的，首字母小写是私有的)。

#### golang标识符命名规则

- Go 语言标识符由 26 个英文字母大小写，0-9 ，_ 组成
- Go 语言标识符不能以数字开头，也就是说允许：下划线或者英文字母开头
- Go 语言标识符严格区分大小写
- Go 语言标识符不能包含空格、@、% 以及 $ 等特殊字符。这一点有区别于java（java变量名允许包含美元符号，允许下划线，数字或者美元符号开头）
- 下划线 _ 本身在 Go 中是一个特殊的标识符，称为 空标识符。可以代表任何其它的标识符，但是它对应的值会被忽略(比如：忽略某个返回值)。所以仅能被作为占位符使用，不能作为标识符使用
- 不能以系统保留关键字作为标识符（一共有25 个）

### golang程序

- 所有的 Go 语言 程序的文件都是以 .go 做为结尾的，即源文件都是以 go结尾
- Go 语言中每个程序文件都必须归属于一个包
- Go语言中程序文件主要提供函数，函数中包含方法
- 对于程序文件的使用，遵循“先引入(import) 再使用”的原则，引用之后即可使用，当然程序文件前提必须是存在的
- Go 应用程序的执行入口是 main() 函数
- Go语言程序中，方法由一条条语句构成，语句语句之间不需要加分号分隔(Go 语言会在每行后自动加分号)
- Go 编译器是逐行编译的，因此有一行就写一条语句，不能把多条语句写在同一行，否则会报错，如果要强行把多行写在同一行，那么可以在每一行语句后加上分号，但是强烈不推荐这样做
- Go 语言定义的变量或者 import 的包如果没有使用到，代码不能编译通过
- Go 语言的大括号都是成对出现的，缺一不可，同时Go 语言中左大括号必须写在上一行代码的后面，不能新起一行，像下面这样：

```go

package main
import "fmt"
//错误写法，golang如果没有再main()后找到大括号，会自动加上一个大括号，所以会报错
func main()
{
	fmt.Println("Hello 嗨客网(www.haicoder.net)")
}

//正确写法
func main(){
	fmt.Println("Hello 嗨客网(www.haicoder.net)")
}
```

- Go 语言应用程序的执行入口是 main() 函数，也就是说一个应用程序必须包含main函数，该应用程序才能被执行。
- 用 + 拼接字符串时，如果要换行，+ 必须留在行末

```go
//正确写法，+号必须紧跟首行末尾
var str string = "hello world " + 
    "hello world"

//错误写法，因为go编译器会在每行末尾自动加上分号
var str string = "hello world"
+ "hello world"
```

- 回车和换行：

```golang
//  \r 是回车符（return），作用是使光标移动到本行的开始位置；
//  \n 是换行符 (newline)，作用是光标垂直向下移动一格
//所以一般用\r\n 来新起一行,且换行后让光标自动回到新行的起始位置
//golang中，回车就是回车，换行就是换行
fmt.Println("hello\nworld!")

//下面这行语句效果：只会打印world!
fmt.Println("hello\rworld!")

//所以记住：要换行，就使用\r\n
```

- golang是静态类型语言，意味着编译期就会对变量类型进行检查，变量需要有明确类型
- int类型默认值为0，float类型默认值为0.0，boolean类型默认为false, String类型默认值为空字符串，指针类型默认值为nil

```txt
bool 	    布尔类型
string 	    字符串类型
int 	    整型
int8 	    整型
int16 	    整型
int32 	    整型
int64 	    整型
uint 	    无符号整型
uint8 	    无符号整型
uint16 	    无符号整型
uint32 	    无符号整型
uint64 	    无符号整型
uintptr 	指针类型
byte 	uint8 的别名
rune 	int32 的别名 代表一个 Unicode 码
float32 	浮点型
float64 	浮点型
complex64 	复数类型
complex128 	复数类型
```

- 变量定义：

```golang
//一般编程语言中，定义变量是： 类型 变量名 = 值
//golang中： var 变量名 类型 = 值

var i int = 10

//函数传参时，也是如此：
func test(str string){
    fmt.Println(str) //Println中P是大写，表示是暴露出来给别的包使用的函数
}
```

- 类型推导:

```golang
//不声明类型直接赋值，会根据值自行判定变量类型
var i = 10 //自行推导出i为int类型
```

- 省略var的情况：

```golang
name := "hello world"
//注意：此时的:不能省略，否则变成赋值了，:=表示声明并赋值，等价于：
var name string
name = "hello world"
```

- 一次声明多个变量:

```golang
//方法一
var n1, n2, n3, n4 int
//方法二
var n1, s1 = 10, "string 1"
//方法三，使用类型推导
n1 s1 := 10, "string1"
```

- 批量声明:

```golang
var (
    name = "Tom"
    age = 30
)
```

- 全局变量声明：

```golang

package main
import "fmt"
//函数体之外声明的变量，被称为全局变量
var globalVariable = 10
func main(){
    fmt.Println(globalVariable)
}
```

- 匿名变量：

```golang
//golang中，一个函数可以返回多个返回值，而如果某个返回值只是声明却不使用，golang运行该程序就会报错。
//当函数的多个返回值，有的返回值不想用的时候，用匿名变量代替，即可避免编译期报错
package main
import "fmt"

func returnMultipleValues(){
    return 0, 11
}

func main(){
    x, y := returnMultipleValues()
    fmt.Println("程序会报错，因为仅仅使用了x，没有使用y",x)
}
```

```golang
package main
import "fmt"

func returnMultipleValues(){
    return 0, 11
}

func main(){
    x, _ := returnMultipleValues()
    fmt.Println("程序不会报错，因为虽然仅仅使用了x，但是使用了匿名变量_",x)
}
```

- 变量的作用域，分为局部作用域和全局作用域，先来说局部作用域：

```golang
//变量是有其作用域的，先说明局部变量的作用域
package main
import "fmt"

func main(){
    //声明局部变量，通过{}声明局部变量域
    {
        name := "Harry Porter"
        fmt.Println("Name = ", name)
    }

    fmt.Println("Hello")
    fmt.Println("执行到我这里会报错，因为我试图访问局部变量name", name)

}

//另外一个局部变量域
func anotherMain(){
    for i :=0 ; i < 3; i++{
        fmt.Print(i)
        fmt.Print("\n")
    }
}

fmt.Println("我不会执行的，也会报错的，因为超出了变量i的作用域",i)
```

- golang全局作用域，全局变量：在函数之外声明的变量,同时分为两类，一类是包内使用的相当于private，其变量名首字母是小写的。第二类是包外使用，变量名首字母大写。
在 Go 语言中，变量、结构体、函数的导出属性是通过它们的首字母大小写来决定的。首字母大写的变量、结构体、或者函数都是导出的，它们可以在整个程序的任何位置，任何包进行访问。
首字母小写的全局变量，结构体、或者函数只能在本包进行访问。

```golang
package main
import "fmt"

//声明了全局变量，即在func之外声明的变量，但是该变量只能在该main函数包内使用，因为首字母是小写n
var name = "Harry Porter"

func main(){
    fmt.Println("我是可以访问到name的，", name)
}
```

```golang
package main
import "fmt"

//声明了main包外都可以使用的全局变量，因为其首字母N是大写的
var Name = "Global Harry Porter"

func main(){
    fmt.Pritnln("Name是可以在所有地方使用到的全局变量，因为其N是大写的")
}
```

- golang常量，表示程序运行期间不会再被修改的量，其数据类型可以是boolean, 数字型(整型，浮点型，复数)和字符串类型

```golang
//声明：const identifier [data type] = value
//可以显式声明
const b string = "abc"
//也可以隐式声明
const anotherB = "abc"
```

- golang的常量也可以用作枚举用途

```golang
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```

- golang常量仅可以使用内置函数对其进行计算，比如len(), cap(), unsafe.SizeOf()等计算。函数必须是内置函数，否则编译会报错

```golang
package main
import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)
func main(){
    println(a, b, c)
}
```

- 由于常量的赋值是一个编译期行为，所以右值不能出现任何需要运行期才能得出结果的表达式，比如试图以如下方式定义常量就会导致编译错误:

```golang
const Home = os.GetEnv("HOME") //os.GetEnv("HOME")明显是运行期才能得到结果，所以编译会报错，右边无法作为值并被赋值给Home常量
```

- 如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式

```golang
//https://studygolang.com/articles/5296

const (       // iota被重设为0
    c0 = iota // c0 == 0
    c1 = iota // c1 == 1
    c2 = iota // c2 == 2
)

const (
    a = 1 << iota // a == 1 (iota在每个const开头被重设为0)
    b = 1 << iota // b == 2
    c = 1 << iota // c == 4
)
```

上面的代码可以被简写为

```golang
const (       // iota被重设为0
    c0 = iota // c0 == 0
    c1        // c1 == 1
    c2        // c2 == 2
)
const (
    a = 1 <<iota // a == 1 (iota在每个const开头被重设为0)
    b            // b == 2
    c            // c == 4
)
```

- Go 语言预定义了这些常量：true、false 和 iota,其中比较特殊的常量：iota，可以被认为是一个可以被编译器修改的常量，遇到const关键字，即被重置为0(即const内部的第一行之前)，const内每增加一行，就是的iota被计数一次

## golang数据类型

- 其中golang的数字类型、布尔型、字符型被称为内置类型，当某个参数类型为golang内置类型时，传递的是副本
- golang派生类型：

```txt
(a) 指针类型（Pointer）
(b) 数组类型
(c) 结构化类型(struct)
(d) Channel 类型
(e) 函数类型
(f) 切片类型
(g) 接口类型（interface）
(h) Map 类型 
```

### golang内置数据类型-数值类型-整型int

- 整型类型有很多：int, int8, int16, int32, int64,uint
- 在32位系统中占4个字节，64位系统中占8个字节
- 整数**默认**推动类型是int类型

### golang内置数据类型-数值类型-浮点数

- 只有float32, float64
- 浮点数类型默认声明为64位

```golang
var f = 3.14
fmt.Printf("%T", f)
//output: float64
```

- 也可以用科学技术法(e或者E都可以)：

```golang
var f = 3.14E2
```

### golang内置数据类型-字符类型

- golang中没有专门的字符类型，如果要存储单个字符，一般用byte保存
- golang中字符串由字节组成，其他语言中字符串是由字符组成，这点是不同的

```golang
var c byte = 'a'
fmt.Println("c = ", c)
//output: 97
```

- 直接输出字符时，输出的不会是字符，而是字符的编码值，如果要输出字符，必须使用格式化输出：

```golang
fmt.Printf("c = %c", c) //使用格式化输出函数Printf
```

- golang中，中文可以是一个字符，这点在其他编程语言中是不可能的，但是中文的编码值明显超出byte的范围，可以用int或者rune定义，**Golang中统一使用UTF-8编码**：

```golang
var c rune = '中'
fmt.Printf("%v %c",c,c)
//output:
```

- 字符串，多行字符串使用`引起来

```golang
var str string = `hello
world`
```

- 字符默认值是""

- 变量之间的转换，golang中没有变量的隐式转换，变量与变量之间必须显式转换：

```golang
var i int = 30
var f float32 = float32(i)
```

- 大数据类型转换成小数据类型(比如int64转换成int32时)，可能会溢出(这点与其他编程语言类似)，但是，溢出不会报错(比较坑了):

```golang
var bigNumber int64 = 99999
var smallNumber int8 = int8(bigNumber) //转换结果：63
```

- 不同类型，无法比较：

```golang
var number1 int64 = 33
var number2 int32 = 33

//下面语句将会是非法的，因为golang判定number1, number2因为是不同类型，所以无法相互比较
if number1 == number2{
    fmt.Println("我将无法被打印")
}
```

### golang内置数据类型-布尔型

- golang布尔型也被称为bool类型
- bool只允许取true, false, true, false本身也是golang的常量
- bool占1个字节，默认值为false
- 用于流程控制和条件判断

```golang
package main
import "fmt"

func main(){
    fmt.Println("Hello")
    var isOk bool
    var isOneline = true
    fmt.Println("isOk:", isOk, "isOnline: ", isOnline)//输出isOk: false, isOnline: true
}
```

## golang引用类型

- 当传递的参数，为一个函数的时候，传递的就是参数其本身，如果是内置类型，传递的是值的副本
- 常见引用类型：切片，map, channel, interface和函数func类型

### golang引用类型-切片

- 英文名为slice
- 切片，可以被理解为数组的引用
- 切片是引用类型，遵守引用类型传递机制：即传递的是其本身，而不是值的副本
- 切片使用和数组类似，遍历、访问切片元素、求切片的长度与数组都一样
- 切片长度可以变化，不像数组是固定的
- 切片是一个可以动态变化的数组

- 创建切片,只需要指定切片类型，不需要指定切片的长度：

#### 数组声明以及数组元素的修改
```golang
// var sliceName []Type: Type是每个元素的类型

package main
import (
    "fmt"
)

func main(){
    //创建了一个有2个元素，且每个元素都是string类型的数组，创建的同时完成初始化
    var slieceString = []string{"Stringa", "Stringb"}
    fmt.Println("slieceString=", slieceString)//output: slieceString= [Stringa Stringb]

    // 注意：数组不是切片，切片是数组的抽象
    // 通过[]索引修改数组
    arr := [5]int{1, 2, 3, 4, 5}
    for i := 0; i < len(arr); i++ {
      arr[i] += 100
    }
    fmt.Println(arr)  // [101 102 103 104 105]
}

```

#### 切片声明以及修改
```golang
// golang slice:
// https://go.dev/blog/slices-intro
// A slice can be created with the built-in function called make, which has the signature
func make([]T, len, cap) []T
// where T stands for the element type of the slice to be created. The make function takes a type, a length, and an optional capacity. When called, make allocates an array and returns a slice that refers to that array.


//使用TypeOf函数，获取切片类型
// 数组的长度不能改变，如果想拼接2个数组，或是获取子数组，需要使用切片。切片是数组的抽象。 切片使用数组作为底层结构。切片包含三个组件：容量，长度和指向底层数组的指针,切片可以随时进行扩展

// 声明切片
slice1 := make([]float32, 0) // 长度为0的切片
slice2 := make([]float32, 3, 5) // [0 0 0] 长度为3容量为5的切片
fmt.Println(len(slice2), cap(slice2)) // 3 5

// 使用切片
slice2 = append(slice2, 1, 2, 3, 4) // [0, 0, 0, 1, 2, 3, 4]
fmt.Println(len(slice2), cap(slice2)) // 7 12
// 子切片 [start, end)
sub1 := slice2[3:] // [1 2 3 4]
sub2 := slice2[:3] // [0 0 0]
sub3 := slice2[1:4] // [0 0 1]
// 合并切片
combined := append(sub1, sub2...) // [1, 2, 3, 4, 0, 0, 0]


// 声明切片时可以为切片设置容量大小，为切片预分配空间。在实际使用的过程中，如果容量不够，切片容量会自动扩展。
// sub2... 是切片解构的写法，将切片解构为 N 个独立的元素。


```

- 查看切片类型，使用TypeOf函数：

```golang
package main
import (
    "fmt"
    "reflect"
)

func main(){
    fmt.Println("使用TypeOf查看切片类型")
    var sliceString = []string{"stringa", "stringb"}
    fmt.Println("sliceString type: ", reflect.TypeOf(sliceString))// output: sliceString type: []string
}
```

- 访问切片元素，和访问数组一样：

```golang
package main
import ("fmt")

func main(){
    var sliceString = []string{"stringa","stringb"}
    fmt.Println("slieceString[0] = ", sliceString[0])
    fmt.Println("slieceString[1] = ", sliceString[1])
}
```

### golang引用类型-map

- 是一个key-value的无序的集合
- 可被称为关联数组或字典
- 可根据指定的key，快速获取value
- key可以是任何可以使用 == 进行比较的数据类型，比如：int, string, bool等
- value可以是任意类型
- 无序，所以每次遍历的顺序可能是不一致的

- 声明，通过var关键字声明：

```golang
//var mapName map[keyType]valueType，声明一个key为keyType类型,value为valueType类型的map,名为mapName

// 仅声明
m1 := make(map[string]int)
// 声明时初始化
m2 := map[string]string{
	"Sam": "Male",
	"Alice": "Female",
}
// 赋值/修改
m1["Tom"] = 18


package main
import (
    "fmt"
)

func main(){
    //map直接赋值使用的是：key:value的形式
    herosMap := map[string]string{
        "hero1":"宋江"
        "hero2":"卢俊义"
        "hero3":"徐达"
    }
    fmt.Println(herosMap) //output: map[hero1:宋江 hero2:卢俊义 hero3:徐达]
}
```

### map返回值为2个
```golang
// https://go.dev/blog/maps
// A two-value assignment tests for the existence of a key:

// In this statement, the first value (i) is assigned the value stored under the key "route". If that key doesn’t exist, i is the value type’s zero value (0). The second value (ok) is a bool that is true if the key exists in the map, and false if not.

i, ok := m["route"]

```


- 获取map类型，依然是TypeOf函数：

```golang
package main
import (
    "fmt"
    "reflect"
)

func main(){
    herosMap := map[string]string{
        "hero1":"宋江"
        "hero2":"卢俊义"
        "hero3":"徐达"
    }
    fmt.Println(reflect.TypeOf(herosMap))// output: map[string]string
}
```

- 访问map元素，有且只能通过key来访问:

```golang
package main
import (
    "fmt"
    "reflect"
)

func main(){
    herosMap := map[string]string{
        "hero1":"宋江"
        "hero2":"卢俊义"
        "hero3":"徐达"
    }
    fmt.Println(herosMap["hero3"])// output: 徐达
}
```

### golang引用类型-channel

- chan是golang语言中核心类型
- 第一次接触chan，直接可以将其看作是管道就行了，既然是管道，就可以是单向管道，也可以是双向管道：https://www.runoob.com/w3cnote/go-channel-intro.html#Channel%E7%B1%BB%E5%9E%8B
- 通过chan就可以并发核心单元(稍后会解释)，就可以发送/接收数据
- 其本质是队列，且线程安全，自带锁的功能
- 操作符：<-，箭头的指向就是数据的流向,没有指定方向，那么channel就是双向，既可以接受数据，也可以发送数据
- 必须是先创建再使用

- channel类型的定义格式：

```golang
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .


```

### golang结构类型

```golang

```

## golang 注释

- 合理注释代码应当占总代码量的1/3
- golang中分为单行注释和多行注释，允许嵌套使用：

```golang
//单行注释
package main
import "fmt"
func main(){
    //fmt.Println("这是单行注释，不会被执行")
    fmt.Println("这里会被执行")
    /* 这也是单行注释 */
}
```

```golang
//块注释
package main
import "fmt"
func main(){
    /*
    * fmt.Println("这里是多行注释，不会被执行")
    */
    fmt.Println("hello world!")

}
```

- 块注释不允许嵌套使用：

```golang
package main
import "fmt"
func main(){
    /*
    *单行注释
     /**
       * 非法的嵌套注释，不允许使用
    */
    fmt.Println("我根本没有机会被执行，因为上面有非法嵌套注释");
}
```

## golang 运算符

### 算术运算符

```txt
+	相加	A + B 输出结果 30
-	相减	A - B 输出结果 -10
*	相乘	A * B 输出结果 200
/	相除	B / A 输出结果 2
%	求余	B % A 输出结果 0
++	自增	A++ 输出结果 11
--	自减	A-- 输出结果 9
```

### 关系运算符

```txt
==	检查两个值是否相等，如果相等返回 True 否则返回 False。	(A == B) 为 False
!=	检查两个值是否不相等，如果不相等返回 True 否则返回 False。	(A != B) 为 True
>	检查左边值是否大于右边值，如果是返回 True 否则返回 False。	(A > B) 为 False
<	检查左边值是否小于右边值，如果是返回 True 否则返回 False。	(A < B) 为 True
>=	检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。	(A >= B) 为 False
<=	检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。	(A <= B) 为 True 
```

### 逻辑运算符

```txt
&&	        逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 	        (A && B) 为 False
||	        逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。	        (A || B) 为 True
!	        逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。	            !(A && B) 为 True 
```

### 位运算符

```txt
&	        按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 	                                        (A & B) 结果为 12, 二进制为 0000 1100
|	        按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或	                                            (A | B) 结果为 61, 二进制为 0011 1101
^	        按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。	        (A ^ B) 结果为 49, 二进制为 0011 0001
<<	        左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 	A << 2 结果为 240 ，二进制为 1111 0000
>>	        右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 	A >> 2 结果为 15 ，二进制为 0000 1111
```

### 赋值运算符

```txt

=	                简单的赋值运算符，将一个表达式的值赋给一个左值	   C = A + B 将 A + B 表达式结果赋值给 C
+=	                相加后再赋值	                                C += A 等于 C = C + A
-=	                相减后再赋值	                                C -= A 等于 C = C - A
*=	                相乘后再赋值	                                C *= A 等于 C = C * A
/=	                相除后再赋值	                                C /= A 等于 C = C / A
%=	                求余后再赋值	                                C %= A 等于 C = C % A
<<=	                左移后赋值 	                                    C <<= 2 等于 C = C << 2
>>=	                右移后赋值 	                                    C >>= 2 等于 C = C >> 2
&=	                按位与后赋值	                                C &= 2 等于 C = C & 2
^=	                按位异或后赋值	                                C ^= 2 等于 C = C ^ 2
|=	                按位或后赋值	                                C |= 2 等于 C = C | 2
```

### 其他运算符

```txt
&	                返回变量存储地址	                &a; 将给出变量的实际地址。
*	                获取指针变量对应的值                *a; 是一个指针变量
```

```golang
package main
import "fmt"
func main(){
    var a int = -4
    var b int32
    var c float32
    var ptr *int

    fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
    fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
    fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

    /*  & 和 * 运算符实例 */
    ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
    fmt.Printf("a 的值为  %d\n", a);
    fmt.Printf("*ptr 为 %d\n", *ptr);


    //总结
    var aa = 100 //int类型，值为100
    var bb *int = &a //只是声明获取int类型指针变量地址，归根结底其还是一个地址，就是指针
    //或者
    var bbb = &a //类型推导，获取bbb作为一个int类型的指针
    //使用，不管是推导方式，还是显式声明的，都需要通过*variable方式
    fmt.Println("获取bb的值:", *bb)
    fmt.Println("获取bbb的值:", *bbb)
}
```

```golang
package main
import (
    "fmt"
    "strings"
    )

func main(){

    fmt.Println("Just input your name")
    var name string
    fmt.Scanln(&name)
    //golang捕捉到了我们输入name之后，摁的ENTER，所以默认的output会有换行动作，为了消除换行，我们需要执行下面语句消除空白
    //事实证明上面的注释所说是错误的，运行second-play/test.go，即可发现，golang 1.16.3中并没有出现换行的情况
    name = strings.TrimSpace(name)
    fmt.Printf("Hello!, %s I'm Golang!", name)

}

```

## golang-结构体

http://c.biancheng.net/view/65.html

结构体，pointer receiver和value receiver的区别：
https://juejin.cn/post/6963476381728702501

https://cyent.github.io/golang/method/method_receiver/

```go
package main

import "fmt"

// Person ...
type Person interface {
    SetAge()
    SetName()
}

// Teacher ...
type Teacher struct {
    Name string
    Age  int
}

// SetAge ...
func (p *Teacher) SetAge() {
    p.Age = 24
}

// SetName ...
func (p *Teacher) SetName() {
    p.Name = "teacher"
}

// Student  ...
type Student struct {
    Name string
    Age  int
}

// SetAge ...
func (p Student) SetAge() {
    p.Age = 16
}

// SetName ...
func (p Student) SetName() {
    p.Name = "student"
}

func main() {
    var p1 Person = &Teacher{}
    p1.SetAge()
    p1.SetAge()
    fmt.Printf("teacher: %+v \n", p1) // teacher: &{Name: Age:24}

    var p2 Person = Student{}
    p2.SetAge()
    p2.SetAge()
    fmt.Printf("Student: %+v \n", p2) // Student: {Name: Age:0}

    var p3 Person = &Student{}
    p3.SetAge()
    p3.SetAge()
    fmt.Printf("Student: %+v \n", p3) // Student: {Name: Age:0}

    var p4 Person = Teacher{}
    p4.SetAge()
    p4.SetAge()
    fmt.Printf("teacher: %+v \n", p4) // cannot use Teacher literal (type Teacher) as type Person in assignment: Teacher does not implement Person (SetAge method has pointer receiver)
}

```

### 结构体定义

- golang中允许通过自定义方式，定义新的类型。，结构体就是这些类型中的一种，且是复合类型
- 结构体是由0个或者多个任意类型的值聚合撑的实体，每个值都可以被称为结构体的成员
- 结构体中的字段拥有自己的类型和值
- 结构体中的字段名称必须唯一
- 结构体中的字段可以是结构体，甚至是字段所在的结构体的类型
- 结构体中的字段，如果是同一类型，，可以写在同一行
- 声明结构体不会分配内存，只是一种内存布局的描述
- 结构体的默认值是nil,结构体没有真正的值

```golang
//结构体的名称在包内不能重复
type <structure-name> struct{
    <field1> <field1-type>//fielda string
    <field2> <field2-type>

    R, G, B byte//同种类型声明在同一行
}
```

## golang-type

http://c.biancheng.net/view/25.html

https://wizardforcel.gitbooks.io/go42/content/content/42_17_type.html

- golang中牵涉到type的有2个地方：类型定义的时候使用到type关键字，类型别名的时候也使用到type关键字
- 是golang1.9之后添加的功能

```golang
//golang1.9之前：
//定义新类型
type byte uint8
type rune int32
//golang1.9之后：
type byte = uint8
type rune = int32
```

### 使用type创建自定义类型

- 既可以使用type关键字定义一个新的结构体
- 也可以使用type关键字，以一个已存在的类型作为基础类型，定义新类型，被称为自定义类型
- 定义过的类型也就拥有了新的特性
- 新类型不会拥有原基础类型所附带的方法

```golang
//声明新类型，以int类型作为基础类型
type newinttype int
//使用新类型newinttype来创建变量
var newValue newinttype

//定义多个自定义类型
type (
    firstOne int
    secondOne float64
    thirdOne string
)

//在golang中，每个值都必须在经过编译之后，属于某个类型，因为其是静态类型语言
//所有的类型的值必须显式说明，显式转换
//类型A的值=类型B(类型A的值)

//值的类型转换:
//定义新类型typeB
type typeB int
//创建名为valueOfTypeB的typeB类型的变量
var valueOfTypeB typeB
//显式转换valueOfTypeA的值给变量valuleOfTypeB
valueOfTypeB = typeB(valueOfTypeA)

```

### 使用type作为类型别名

- 为某数据类型起一个别名
- 别名类型与原类型可被视作同一种类型
- 别名类型拥有原类型附带的所有方法

```golang
package main

import "fmt"

type A struct{
    Face int
}

//类型别名
type anotherA = A

//A的方法
func (a A) speak() {
    fmt.Println("A speak aaa", a.Face)
}

func main(){
    var ainstance A = A{Face: 9}
    ainstance.speak()

    var anotherAinstance anotherA = anotherA{Face: 9}
    anotherAinstance.speak()
}

//output
// A speak aaa 9
// A speak aaa 9
```

## Explicit Conversion

https://golang.org/ref/spec#Conversions

```golang
//Golang conventions
// https://golang.org/ref/spec#Conversions


newDataTypeVariable = T(oldDataTypeVariable)

// Golang does not support implicit type conversion because of its robust type system.
// Explicit Type Conversion
//Explicit type conversion is special programming instruction that specifies what data type to treat a variable as (or an intermediate calculation result) in a given expression. For example, casting will ignore “extra” information (but never adds information to the type being cast).

// See the following code snippets of explicit type casting in Go.
var badboys int = 1921

// explicit type conversion
var badboys2 float64 = float64(badboys)
var badboys3 int64 = int64(badboys)
var badboys4 uint = uint(badboys)

// 也可以看以下例子：
// https://go101.org/article/value-conversions-assignments-and-comparisons.html
//https://notes.shichao.io/gopl/ch7/
// *bytes.Buffer must satisfy io.Writer
var _ io.Writer = (*bytes.Buffer)(nil)
// 将nil转为io.Writer类型，如果bytes.Buffer没有实现Writer接口的方法，转换将失败报错

```

### 空接口
如果定义了一个没有任何方法的空接口，那么该接口可以表示任意类型
```golang
func main() {
	m := make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 18
	m["scores"] = [3]int{98, 99, 85}
	fmt.Println(m) // map[age:18 name:Tom scores:[98 99 85]]
}

```

## golang switch

```golang
type Gender int8
const (
	MALE   Gender = 1
	FEMALE Gender = 2
)

gender := MALE

switch gender {
case FEMALE:
	fmt.Println("female")
case MALE:
	fmt.Println("male")
default:
	fmt.Println("unknown")
}
// male

```

在这里，使用了type 关键字定义了一个新的类型 Gender。

使用 const 定义了 MALE 和 FEMALE 2 个常量，Go 语言中没有枚举(enum)的概念，一般可以用常量的方式来模拟枚举。

和其他语言不同的地方在于，Go 语言的 switch 不需要 break，匹配到某个 case，执行完该 case 定义的行为后，默认不会继续往下执行。如果

需要继续往下执行，需要使用 fallthrough，例如：

```golang
switch gender {
case FEMALE:
	fmt.Println("female")
	fallthrough
case MALE:
	fmt.Println("male")
	fallthrough
default:
	fmt.Println("unknown")
}
// 输出结果
// male
// unknown
```

## golang for循环

直接访问：https://tour.golang.org/flowcontrol/1


- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

for循环中的3条件彼此以分号分割，其中init和post可以省略

```golang
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

```

如果省略了condition，那么for循环将永远执行



```golang
// for循环和java等没有区别
sum := 0
for i := 0; i < 10; i++ {
	if sum > 50 {
		break
	}
	sum += i
}
```

如果是对数组(arr)、切片(slice)、字典(map)， 使用 for range 遍历:
```golang
// 如果是对数组进行遍历，需要借助索引
nums := []int{10, 20, 30, 40}
for i, num := range nums {
	fmt.Println(i, num)
}
// 0 10
// 1 20
// 2 30
// 3 40

// 如果是对slice进行遍历，需要借助key, value
m2 := map[string]string{
	"Sam":   "Male",
	"Alice": "Female",
}

for key, value := range m2 {
	fmt.Println(key, value)
}
// Sam Male
// Alice Female
```

## Golang Error Handling
如果函数实现过程中，如果出现不能处理的错误，可以返回给调用者处理。比如我们调用标准库函数os.Open读取文件，os.Open 有2个返回值，第一个是 *File，第二个是 error， 如果调用成功，error 的值是 nil，如果调用失败，例如文件不存在，我们可以通过 error 知道具体的错误信息。

```golang
import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err)
	}
}

// open filename.txt: no such file or directory
```

**唯一亮点就是：可以通过 errorw.New 返回自定义的错误**

```golang
import (
	"errors"
	"fmt"
)

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}

// if的另外简洁写法
func main() {
	if err := hello(""); err != nil {
		fmt.Println(err)
	}
}
// error: name is null
```

### Panic

https://go.dev/blog/defer-panic-and-recover

类似于java中的Unchecked Exception. error 往往是能预知的错误，但是也可能出现一些不可预知的错误，例如数组越界，这种错误可能会导致程序非正常退出，在 Go 语言中称之为 panic。

Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.

Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.

```golang
func get(index int) int {
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func main() {
	fmt.Println(get(5))
	fmt.Println("finished")
}
```

```shell
# 以下是运行结果
$ go run .
panic: runtime error: index out of range [5] with length 3
goroutine 1 [running]:
exit status 2
```

```golang
// 以下是一个完整的例子
package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}


// 运行结果：
// Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.



// If we remove the deferred function from f the panic is not recovered and reaches the top of the goroutine’s call stack, terminating the program. This modified program will output:
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
panic: 4

panic PC=0x2a9cd8
[stack trace omitted]
```

### Golang Try Catch

golang defer, recover and panic: https://go.dev/blog/defer-panic-and-recover

在golang中是没有try...catch...的，但是golang提供了另外一种方式：defer 和 recover

代码中使用defer后面跟一个function或者method调用，该defer会将对于某个function或者method的调用推送到一个执行列表中，defer所在的外部的function在return之前会执行该defer对应的function call

**注意：**只要定义了defer语句，那么defer语句就一定会执行

```golang
func get(index int)(ret int){
  // 使用defer定义异常处理的函数，协程退出之前，会执行完defer挂载的任务。也就是外部函数在return之前会执行defer对应的调用。只要get函数某行代码触发了panic，就会把控制权交给defer声明的异常处理函数
  defer func(){
    // 使用recover，如果recover不为nil，表明发生了异常
    if r := recover();r!=nil{
      fmt.Println("Some error happened", r)
      // 使用recover将程序恢复正常，并将ret设置为-1，如果不设置ret，默认返回值为0
      ret = -1
    } 
  }()

  arr := [3]int{2,3,4}
  return arr[index]
}

func main(){
  fmt.Println(get(6))
  fmt.Println("finished")
}

// 运行结果：
// $ go run .
// Some error happened! runtime error: index out of range [5] with length 3
// -1
// finished
```

defer可以用于以下场景：
- 异常catch
- clean-up functions

```golang
// https://go.dev/blog/defer-panic-and-recover
// 如果没有defer的情况
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
// 如果在下面的os.Create方法发生异常或者错误，那么dst.Close, src.Close将不会被执行
    dst, err := os.Create(dstName)
    if err != nil {
        return
    }

    written, err = io.Copy(dst, src)
    dst.Close()
    src.Close()
    return
}

// 使用defer的情况：
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    // 一旦os.Open发生错误，在return err之前，执行src.Close()
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    // 同理，在第二个return之前
    defer dst.Close()

    return io.Copy(dst, src)
}

```

由上得知，defer的声明一定是在return之后的

### defer rule 1

A deferred function’s arguments are evaluated when the defer statement is evaluated.

```golang
package main

import "fmt"

func a() (i int){
	test := 3
  // 注意：defer语句打印的test的值不会是4，而是3，这里只会打印test的初始值
	defer fmt.Println("defer test:",test)
	test = 4
	fmt.Println("This will print before defer, and test:", test)
	return test
}

func main() {

	fmt.Println("Hello, 世界")
	a()
	fmt.Println("This print at last")
	
}

// 执行结果：
// Hello, 世界
// This will print before defer, and test: 4
// defer test: 3
// This print at last
```

### defer rule 2

Deferred function calls are executed in Last In First Out order after the surrounding function returns.

**永远记住：defer是在外部的function return之后才会被执行**, 后进先出原则

```golang
package main

import "fmt"

func b() (a int){
	for i:=0; i<4; i++ {
		defer fmt.Println(i);
	}
	fmt.Println("The for loop ends up.")
	return 1
}
func main() {

	fmt.Println("Hello, 世界")
	b()
	
}

// Hello, 世界
// The for loop ends up.
// 3
// 2
// 1
// 0
```

## resoruces

https://cloud.tencent.com/developer/article/1386519

http://c.biancheng.net/view/23.html

https://studygolang.com/articles/5296

https://haicoder.net/golang/golang-mutex.html

https://www.digitalocean.com/community/tutorial_series/how-to-code-in-go

https://assets.digitalocean.com/books/how-to-code-in-go.pdf

https://wizardforcel.gitbooks.io/go42/content/content/42_17_type.html

https://gobuffalo.io/en/docs/getting-started/installation


https://go.dev/blog/slices-intro
https://pgk.go.dev
https://golang.org/doc/effective_go