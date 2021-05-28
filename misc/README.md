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

## golang debugging in vscode

```sh
# https://github.com/golang/vscode-go/blob/master/docs/debugging.md
# https://www.thegreatcodeadventure.com/debugging-a-go-web-app-with-vscode-and-delve/
# https://www.digitalocean.com/community/tutorials/debugging-go-code-with-visual-studio-code
# https://github.com/microsoft/vscode-go/issues/3166
```