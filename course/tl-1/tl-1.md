# golang tutorial

牢牢记住：

https://go.dev/doc/

以及

https://go.dev/tour/Welcome/1

以及[go指南首页-中文版-包含教程目录](https://tour.go-zh.org/list)

[go指南-线上编辑器](https://go-zh.org/)

或英文版：https://go.dev/tour/basics/3

built-in go package: https://pkg.go.dev/builtin

## 下载安装Golang package

https://go.dev/doc/install

唯一要注意的是配置环境变量

```sh
# linux ubuntu环境：https://golang.org/doc/install?download=go1.16.4.linux-amd64.tar.gz

wget https://golang.org/dl/go1.17.8.linux-amd64.tar.gz

rm -rf /usr/local/go && tar -C /usr/local -xzvf go1.17.8.linux-amd64.tar.gz

# 将下面内容添加至~/.bashrc(注意root目录下添加，或者自己账号下添加，或/etc/profile)
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/.go

# create golang directory
mkdir ~/.go

# 生效
source ~/.bashrc
# source /etc/profile

# 成功后运行：
go version

```

## first hello world

https://go.dev/doc/tutorial/getting-started

```sh
# 实例参照hello/目录内容
 mkdir hello
 cd hello/
root@marco-VirtualBox:go_workspace/hello# ll
total 8
drwxr-xr-x 2 root root 4096 3月  12 21:55 ./
drwxr-xr-x 3 root root 4096 3月  12 21:55 ../
root@marco-VirtualBox:go_workspace/hello# go mod init learngo/hello
go: creating new go.mod: module learngo/hello
root@marco-VirtualBox: go_workspace/hello# ll
total 12
drwxr-xr-x 2 root root 4096 3月  12 21:55 ./
drwxr-xr-x 3 root root 4096 3月  12 21:55 ../
-rw-r--r-- 1 root root   30 3月  12 21:55 go.mod
root@marco-VirtualBox:go_workspace/hello# cat go.mod
module learngo/hello

go 1.17
root@marco-VirtualBox:go_workspace/hello#


root@marco-VirtualBox:go_workspace/hello# vi hello.go
root@marco-VirtualBox:go_workspace/hello# go run .
Sammy says Hello!
root@marco-VirtualBox:go_workspace/hello#

```

## 访问external package

- 查询[pkg.go.dev](https://pkg.go.dev/search?q=quote)查询package,查看其提供的API
- import "rsc.io/quote"

在国内的话需要设置代理，否则会出现go mod tidy拉取不了package的情况

```sh
export GO111MODULE=on
export GOPROXY=https://goproxy.io
```

注意hello目录下，func main只能出现在一个xx.go文件中，不能同时有多个main方法存在，否则会找不到入口
