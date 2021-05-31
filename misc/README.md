# 杂项

## Type Declaration

https://golang.org/ref/spec#Type_declarations

http://c.biancheng.net/view/25.html

## online console

https://tour.golang.org/basics/4

## cobra

https://github.com/spf13/cobra

```sh
go get -u github.com/spf13/cobra
```

### Debug for cobra

```sh
# https://stackoverflow.com/questions/57290306/debug-file-other-than-main-go-in-vs-code
# https://github.com/golang/vscode-go/blob/master/docs/debugging.md
```

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
            "program": "${workspaceFolder}",
            "env": {},
            "args": [
                "env", // the subcommand of `dxxx env login`
                "login", // the subcommand of `dxxx env login`
            ]
        }
    ]
}
```

### tutorial

```sh
# Use cobra to develop CLI: https://towardsdatascience.com/how-to-create-a-cli-in-golang-with-cobra-d729641c7177
```

## Json Marshal in golang

```sh
https://pkg.go.dev/encoding/json#Marshal


```

## golang debugging in vscode

```sh
# https://github.com/golang/vscode-go/blob/master/docs/debugging.md
# https://www.thegreatcodeadventure.com/debugging-a-go-web-app-with-vscode-and-delve/
# https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code
# https://github.com/microsoft/vscode-go/issues/3166
```

## Set default value in golang for struct

```sh
https://www.geeksforgeeks.org/how-to-assign-default-value-for-struct-field-in-golang/

# struct with json
https://eli.thegreenplace.net/2020/optional-json-fields-in-go/
# field tags in struct
https://pkg.go.dev/encoding/json#Marshal
https://blog.kalan.dev/golang-default-value/

```

## rest api via golang 

```sh
# https://medium.com/@matryer/how-i-write-go-http-services-after-seven-years-37c208122831
# https://tutorialedge.net/golang/creating-restful-api-with-golang/
# https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html

# convert the rest api response to golang struct
# https://blog.alexellis.io/golang-json-api-client/
# write unit test:
# https://blog.alexellis.io/golang-writing-unit-tests/

# json with golang
# https://tutorialedge.net/golang/parsing-json-with-golang/
```

## testing framework

```sh
# https://pkg.go.dev/github.com/onsi/ginkgo@v1.8.0

```

## map interface in golang

```sh
# https://bitfieldconsulting.com/golang/map-string-interface
```

## Parse json file

```sh
# https://www.golangprograms.com/golang-read-json-file-into-struct.html
# https://tutorialedge.net/golang/parsing-json-with-golang/
# https://stackoverflow.com/questions/16681003/how-do-i-parse-a-json-file-into-a-struct-with-go
# https://yourbasic.org/golang/json-example/
```

## defer command in golang

```sh
# https://gobyexample.com/defer
```

## range command

```sh
# https://gobyexample.com/range
```

## empty interface

```sh
# https://stackoverflow.com/questions/59976812/empty-interfaces-in-golang
# https://medium.com/a-journey-with-go/go-understand-the-empty-interface-2d9fc1e5ec72

# golang interfaces
# https://research.swtch.com/interfaces
```