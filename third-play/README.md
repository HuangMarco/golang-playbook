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

