# Go Project related

## 环境搭建

### GOPATH配置

```sh
# 新建目录bin, src
mkdir -p $HOME/go/{bin, src}

# output
#└── $HOME
#   └── go
#       ├── bin
#       └── src

# $GOPATH告诉golang compiler到哪里去家在third party source code
# 虽然现在不需要了，但是很多第三方库依然依赖这个环境变量
export GOPATH=$HOME/go
# 当golang编译和安装tools的时候，会把这些tools放在$GOPATH/bin目录下
# 配置了这个之后，你可以在任何地方运行代码了
export PATH=$PATH:$GOPATH/bin
# 配置了下面这行之后，即可以在任何地方运行所有golang tools
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin

source ~/.zshrc
```

```golang
//配置好了$GOPATH之后，执行下面语句查看本地golang version
go version 
// 拉取https://github.com/digitalocean/godo仓库，即可发现在配置好的$GOPATH目录下多出目录pkg,在其中存放所有拉取的third party golang tools
//在golang较早版本的时候，执行下面的语句会将第三方库拉取到$GOPATH/go/src目录，新版本(大约golang1.8)之后，会将第三方库放置在pkg目录
go get github.com/digitalocean/godo

```

## vscode中golang相关

```sh

# 执行部分golang插件安装之后，会在$GOPATH/bin目录下发现这些可执行文件，都是golang的tools
-rwxr-xr-x  1   staff   3301024 May 25 10:39 go-outline
-rwxr-xr-x  1   staff   8538816 May 25 10:38 gocode
-rwxr-xr-x  1   staff   8321648 May 25 10:40 gocode-gomod
-rwxr-xr-x  1   staff   4373440 May 25 10:28 gopkgs
-rwxr-xr-x  1   staff  22370064 May 25 10:40 gopls

# mkdir  %GOPATH%\\src\\golang.org\\x
# git clone https://github.com/golang/tools.git %GOPATH%\\src\\golang.org\\x\\tools

# 配置好$GOPATH以及相关环境变量之后，执行go get会自行相对应目录，以存放golang tools

go get -v github.com/mdempsky/gocode
go get -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/acroca/go-symbols
go get -v golang.org/x/tools/cmd/guru
go get -v golang.org/x/tools/cmd/gorename
go get -v github.com/derekparker/delve/cmd/dlv
go get -v github.com/stamblerre/gocode
go get -v github.com/rogpeppe/godef
go get -v github.com/ianthehat/godef
go get -v github.com/sqs/goreturns
go get -v github.com/golang/lint%
git clone https://github.com/golang/lint.git %GOPATH%\\src\\golang.org\\x\\lint

go build -o %GOPATH%\\bin\\gocode.exe github.com/mdempsky/gocode
go build -o %GOPATH%\\bin\\gopkgs.exe github.com/uudashr/gopkgs/cmd/gopkgs
go build -o %GOPATH%\\bin\\go-outline.exe github.com/ramya-rao-a/go-outline
go build -o %GOPATH%\\bin\\go-symbols.exe github.com/acroca/go-symbols
go build -o %GOPATH%\\bin\\guru.exe golang.org/x/tools/cmd/guru
go build -o %GOPATH%\\bin\\gorename.exe golang.org/x/tools/cmd/gorename
go build -o %GOPATH%\\bin\\dlv.exe github.com/derekparker/delve/cmd/dlv
go build -o %GOPATH%\\bin\\gocode-gomod.exe github.com/stamblerre/gocode
go build -o %GOPATH%\\bin\\godef.exe github.com/rogpeppe/godef
go build -o %GOPATH%\\bin\\godef-gomod.exe github.com/ianthehat/godef
go build -o %GOPATH%\\bin\\goreturns.exe github.com/sqs/goreturns
go build -o %GOPATH%\\bin\\golint.exe golang.org/x/lint/golint

pause
```
