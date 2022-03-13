# 指针+struct

https://tour.go-zh.org/moretypes/1

## array and slice

官方的[tour](https://tour.go-zh.org/moretypes/11)是非常迷惑人的，关于cap(s)，查看https://gosamples.dev/capacity-and-length/

```go
//创建一个数组{1,2,3,4,5,6}作为underlying array,同时基于该array创建一个slice
s := []int{1,2,3,4,5,6}

start := 0
end := 5
s := s[start:end]

//cap(s) == underlyingArray.length - start,与end无关
//len(s) == end - start
```

## 设置IDEA GOPATH

https://www.jetbrains.com/help/idea/configuring-goroot-and-gopath.html#gopath

GOROOT: https://www.jetbrains.com/help/idea/configuring-goroot-and-gopath.html

IDEA使用Golang: https://www.jetbrains.com/help/idea/go-plugin.html

参照目录下的几个pdf，足够配置好intellij idea拉取golang package
