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
```

## deep source analysis

```sh
# 自动帮助分析检查代码-通常与git hub repo集成使用
# https://deepsource.io/
```

## debug file

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

