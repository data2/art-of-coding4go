# art-of-coding4go
go 语言的小实践

## golang是什么

Go VS Golang由于 Google 在注册 Go 的官网域名时，go.org已被迪士尼抢注，golang.org才得以“上位”，这也就产生了很多人误以为Golang就是其正式名称的误会，而事实却恰恰相反，我们可以认为Golang只是Go的绰号。

## go语言的发展

![image](https://user-images.githubusercontent.com/13504729/130402670-68428d86-655b-425e-9107-c18e586e98de.png)

## Go 语言的发展目标

Go 语言的主要目标是将**静态语言的安全性和高效性与动态语言的易开发性进行有机结合**，达到完美平衡。因此，**Go 语言是一门类型安全和内存安全的编程语言**。虽然 Go 语言中仍有指针的存在，但并不允许进行指针运算。
 
Go 语言的另一个目标是**对于网络通信、并发和并行编程的极佳支持，从而更好地利用大量的分布式和多核的计算机**。设计者**通过 goroutine 这种轻量级线程的概念来实现这个目标，然后通过 channel 来实现各个 goroutine 之间的通信**。他们实现了分段栈增长和 goroutine 在线程基础上多路复用技术的自动化。
 
Go 语言中另一个非常重要的特性就是它的**构建速度（编译和链接到机器代码的速度）**，一般情况下构建一个程序的时间只需要数百毫秒到几秒。
Go 语言在执行速度方面也可以与 C/C++ 相提并论。

## GO并发 

每个Go程序至少拥有一个：main **gotoutine**，当程序开始时会自动创建并启动。在几乎所有Go程序中，你都可能会发现自己迟早加入到一个gotoutine中

Go遵循称为**fork-join模型的并发模型**.fork这个词指的是在程序中的任何一点，它都可以将一个子执行的分支分离出来，以便与其父代同时运行。join这个词指的是这样一个事实，即在将来的某个时候，这些并发的执行分支将重新组合在一起。子分支重新加入的地方称为连接点。这里有一个图形表示来帮助你理解它：

![image](https://user-images.githubusercontent.com/13504729/130546277-13c063d0-fe4b-4892-9f64-c8f3363e7d06.png)

**go关键字为Go程序实现了fork**，fork的执行者是goroutine

``` go
sayHello := func() {
	fmt.Println("hello")
}
go sayHello()
// continue doing other things

或者使用匿名函数

go func(){
 fmt.Println("hello")
}()
```
