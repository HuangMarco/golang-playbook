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

### 整型int

- 整型类型有很多：int, int8, int16, int32, int64,uint
- 在32位系统中占4个字节，64位系统中占8个字节
- 整数**默认**推动类型是int类型

### 浮点数

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

### 字符类型

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

### golang 注释

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


## resoruces

https://cloud.tencent.com/developer/article/1386519

http://c.biancheng.net/view/23.html

https://studygolang.com/articles/5296
