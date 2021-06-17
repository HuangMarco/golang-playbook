# Third Play

## 根据json自动生成golang struct

```golang
// https://mholt.github.io/json-to-go/
```

## downloaded packages not in src but pkg

```sh
export GO111MODULE=off
# 从1.15开始，默认设置了GO111MODULE=on,也就是默认情况下，所有的第三方依赖都会存放到pkg目录
# see https://github.com/golang/go/wiki/Modules for the reason

go env -w GO111MODULE=auto
# do below to execute from package for latest go version
export GO111MODULE="auto"

# show all golang environment variables
go env
```

## validate json

```golang
// https://pkg.go.dev/github.com/valyala/fastjson#Validate
// https://github.com/asaskevich/govalidator
//更多详情见同目录：json-related.go

//注意事项：
//Struct element首字母必须为大写，如此json Encoder/Decoder才能使用这些element.
//json Encoder/Decoder不会使用没有Export过的struct element
```

## format golang

```sh
go fmt <your-golang-program>
```

## map[string]interface {}

部分转载于：https://bitfieldconsulting.com/golang

- [map string interface类型](https://bitfieldconsulting.com/golang/map-string-interface)是一个非常有用的类型，更多用法参照同目录`json-related.go`
- golang中，通过hash table方式来实现Maps,参照[Go maps in action](https://blog.golang.org/maps)
- 通过`m = make(map[string]int)`方式来初始化一个Map type。通过`var m map[string]int`方式声明的m初始化的值是nil，这种方式只是声明一个并没有任何赋值，尝试向一个nil的map塞值会导致报错,make命令会初始化一个hash map数据结构，同时分配内存，然后返回一个空的map(只是空而不是nil)。该过程是在runtime中完成的，而不是golang语言本身完成的
- 注意：在golang中，golang语言不会通过interface{}来实现map，主要是：在编译期间对map声明以及赋值语句进行rewriting，rewriting之后的代码调用的是golang runtime代码，如下，同时注意每当有一个新的map声明，就会复制一份golang runtime的mapaccess(m,k,v)代码，再有新的map声明，就会又复制一份

```golang
v := m["key"]     → runtime.mapaccess1(m, ”key", &v)
v, ok := m["key"] → runtime.mapaccess2(m, ”key”, &v, &ok)
m["key"] = 9001   → runtime.mapinsert(m, ”key", 9001)
delete(m, "key")  → runtime.mapdelete(m, “key”)
```

```golang
func mapaccess1(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer
//key - 指向key的pointer
//h - 指向runtime.hmap结构，该结构存放bucket和其他的自身的参数
// t - 指向用户的maptype

type maptype struct {
       typ           _type
       key         *_type
       elem        *_type
        bucket        *_type // internal type representing a hash bucket 
        hmap          *_type // internal type representing a hmap
        keysize       uint8  // size of key slot
        indirectkey   bool   // store ptr to key instead of key itself
       valuesize     uint8  // size of value slot
      indirectvalue bool   // store ptr to value instead of value itself
       bucketsize    uint16 // size of bucket
       reflexivekey  bool   // true if k==k for all keys
      needkeyupdate bool   // true if we need to update key on overwrite 
}
```
- golang中,map并不是线程安全的，如果从不同的goroutines访问map，可能会导致有多个data race产生，并导致crash，使用[race detector](https://blog.golang.org/race-detector)来避免此类问题
- 在并发环境中，通过使用mutex比如`sync.Map`来避免多个进程同时访问map的时候出现问题。但是最好的解决办法是，在代码中避免concurrent access到同一个map。经典名言"Don't communicate by sharing memory, share memory by communicating",[vedio](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=2m48s)
- 当你把map当成存放数据的box的时候,map就不是pointer，当你在map中存放指向到其他类型的时候，map就是pointer
- 复制map[string]interface{}只能通过for循环的方式，获取其中的每个key,value。或者通过转换为[]byte，再unmarshal的方式来获取值
- map[string]interface转换为yaml:

```golang
import "gopkg.in/yaml.v2"
...
m := map[string]bool{
  "Go is awesome":             true,
  "I drink and I know things": true,
  "PI IS EXACTLY THREE!":      false,
}
data, err := yaml.Marshal(m)
if err != nil {
  log.Fatal(err)
}
fmt.Printf("%s\n", data)
// Go is awesome: true
// I drink and I know things: true
// PI IS EXACTLY THREE!: false
```



## race detector

https://blog.golang.org/race-detector

## deep source analysis - git hub深度代码分析与错误检查

```sh
# 自动帮助分析检查代码-通常与git hub repo集成使用
# https://deepsource.io/
```

## debug file - 调试代码

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "/Users/i323691/work_dir/training/go-training/golang-playbook/third-play/json-related.go",
            "args": ["go", "run","${file}"],
            "dlvLoadConfig": {
                "followPointers": true,
                "maxVariableRecurse": 1,
                "maxStringLen": 400,
                "maxArrayValues": 64,
                "maxStructFields": -1
            }
        }
    ]
}
```

or

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${file}",
            "args": ["go", "run","${file}"],
            "dlvLoadConfig": {
                "followPointers": true,
                "maxVariableRecurse": 1,
                "maxStringLen": 400,
                "maxArrayValues": 64,
                "maxStructFields": -1
            }
        }
    ]
}
```

注意：`${file}`即是你当前vscode中打开的文件，也就是说要debug哪个文件，就保持该文件在vscode当前窗口被打开即可

https://code.visualstudio.com/docs/editor/variables-reference

## interface {}的用法

- 假设你需要根据用户输入打印一段信息，但是用户输入的类型你不知晓，不知道是int还是string，此时可以使用interface{}，interface{}可以代表任何
- func fmt.Println(a ...interface{}) (n int, err error)
- 假设你提前知晓过来的data的element都是string类型，但是element对应的value类型未知，此时可以使用`map[string]interface{}`

```json
{
   "name":"John",
   "age":29,
   "hobbies":[
      "martial arts",
      "breakfast foods",
      "piano"
   ]
}
```

此时可以使用:

```golang
type Person struct {
    Name string `json:"name"`
    Age string `json:"age"`
    Hobbies []string `json:"hobbies"`
}
```

## Positional Arguments

```txt
所谓positional argument位置参数，是指用相对位置指代参数。关键字参数（keyword argument），见名知意使用关键字指代参数。位置参数或者按顺序传递参数，或者使用名字，自然使用名字时，对顺序没有要求。

A positional argument is a name that is not followed by an equal assign（=） and default value.

A keyword argument is followed by an equal sign and an expression that gives its default value.

所以positional arguments有点类似于类型推导

```

### cobra positional arguments

```sh
# https://pkg.go.dev/github.com/spf13/cobra#readme-positional-and-custom-arguments
# https://www.usenix.org/system/files/login/articles/login_summer18_09_mceniry.pdf
# Best tutorial for cobra I believe
# https://www.bookstack.cn/read/cobra/spilt.4.spilt.4.README.md
```

### git

```sh
# compare file in two branches differences
git diff mybranch..master -- <your-file-path>
# compare the local file difference with remote master branch
git diff master -- <your-file-path>
```

## 操作文件

```golang
// https://github.com/spf13/afero
// use golang module "path/filepath"
// use golang module "github.com/spf13/afero"
```

## golang testing

```sh
# https://golang.org/pkg/testing/
# https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter09/09.1.html
# Ginkgo is a BDD-style testing framework for Golang
# https://pkg.go.dev/github.com/onsi/ginkgo#pkg-overview
# https://onsi.github.io/ginkgo/

# https://medium.com/boldly-going/unit-testing-in-go-with-ginkgo-part-1-ce6ff06eb17f
```

## Kubernetes

### Post-Forward

```sh
# https://kubernetes.io/zh/docs/tasks/access-application-cluster/port-forward-access-application-cluster/
```

## print logs

```sh
# github.com/sirupsen/logrus
```


```
package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.wdf.sap.corp/hcpperf/dynatrace-go/afsutil"
	. "github.wdf.sap.corp/hcpperf/dynatrace-go/command"
	"github.wdf.sap.corp/hcpperf/dynatrace-go/dynatrace/environment"
)

var _ = Describe("EnvironmentGetObjectSettings", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
		afsutil.UseVirtualFs()
		_ = environment.AddTarget("default", server.URL()+"/e/12345678-1234-1234-1234-123456789012", "123456789012345678901", false)
		_ = environment.SetActiveTarget("default")
	})

	AfterEach(func() {
		server.Close()
		_ = environment.RemoveTarget("default")
	})

	Context("when no arguments are provided", func() {
		It("should return an error", func() {
			err, uiOut, cobraOut := execCmd(CmdEnvGetObjectSettings(), []string{})

			Expect(err).To(HaveOccurred())
			Expect(uiOut).To(BeEmpty())
			Expect(cobraOut).ToNot(BeEmpty())
			Expect(cobraOut).To(HavePrefix("Error: accepts 2 arg(s), received 0"))
		})
	})

	Context("when there is no object settings configuration", func() {
		It("should return response with empty items and totalCount as 0 and pageSize as 100", func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/e/12345678-1234-1234-1234-123456789012/api/v2/settings/objects"),
					ghttp.RespondWith(200, []byte(`
{"items": [],"totalCount": 0, "pageSize": 100}`)),
				),
			)
			_, uiOut, cobraOut := execCmd(CmdEnvGetObjectSettings(), []string{"environment", "builtin:container.monitoring-rule"})
			Expect(uiOut).To(Equal(`[map[items:[] pageSize:100 totalCount:0]]`))
			Expect(cobraOut).To(BeEmpty())
		})
	})

})

```

## map[string]interface 

```sh
# https://bitfieldconsulting.com/golang/map-string-interface
```

## check is file or directory

```golang
package main

import (
    "fmt"
    "os"
)

func main() {
    name := "FileOrDir"
    fi, err := os.Stat(name)
    if err != nil {
        fmt.Println(err)
        return
    }
    switch mode := fi.Mode(); {
    case mode.IsDir():
        // do directory stuff
        fmt.Println("directory")
    case mode.IsRegular():
        // do file stuff
        fmt.Println("file")
    }
}
```

## array vs slice

```sh
# https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html
# https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter02/02.2.html
```

## []byte convert to string

```sh
# https://segmentfault.com/a/1190000037679588
# https://www.dotnetperls.com/json-go
```

## github review

```sh
# https://www.freecodecamp.org/news/what-do-cryptic-github-comments-mean-9c1912bcc0a4/
```

## goroutines

https://www.golangprograms.com/how-to-use-waitgroup-to-delay-execution-of-the-main-function-until-after-all-goroutines-are-complete.html

## http client

http://dlintw.github.io/gobyexample/public/http-client.html
