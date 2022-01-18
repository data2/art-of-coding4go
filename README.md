# art-of-coding4go
go 语言的小实践

# golang是什么

Go VS Golang由于 Google 在注册 Go 的官网域名时，go.org已被迪士尼抢注，golang.org才得以“上位”，这也就产生了很多人误以为Golang就是其正式名称的误会，而事实却恰恰相反，我们可以认为Golang只是Go的绰号。

# go语言的发展

![image](https://user-images.githubusercontent.com/13504729/130402670-68428d86-655b-425e-9107-c18e586e98de.png)

# Go 语言的发展目标

Go 语言的主要目标是将**静态语言的安全性和高效性与动态语言的易开发性进行有机结合**，达到完美平衡。因此，**Go 语言是一门类型安全和内存安全的编程语言**。虽然 Go 语言中仍有指针的存在，但并不允许进行指针运算。
 
Go 语言的另一个目标是**对于网络通信、并发和并行编程的极佳支持，从而更好地利用大量的分布式和多核的计算机**。设计者**通过 goroutine 这种轻量级线程的概念来实现这个目标，然后通过 channel 来实现各个 goroutine 之间的通信**。他们实现了分段栈增长和 goroutine 在线程基础上多路复用技术的自动化。
 
Go 语言中另一个非常重要的特性就是它的**构建速度（编译和链接到机器代码的速度）**，一般情况下构建一个程序的时间只需要数百毫秒到几秒。
Go 语言在执行速度方面也可以与 C/C++ 相提并论。

# GO并发 

每个Go程序至少拥有一个：main **goroutine**，当程序开始时会自动创建并启动。在几乎所有Go程序中，你都可能会发现自己迟早加入到一个gotoutine中

Go遵循称为**fork-join模型的并发模型**.fork这个词指的是在程序中的任何一点，它都可以将一个子执行的分支分离出来，以便与其父代同时运行。join这个词指的是这样一个事实，即在将来的某个时候，这些并发的执行分支将重新组合在一起。子分支重新加入的地方称为连接点。这里有一个图形表示来帮助你理解它：

![image](https://user-images.githubusercontent.com/13504729/130546277-13c063d0-fe4b-4892-9f64-c8f3363e7d06.png)

**Goroutine是实际并发执行的实体**，它底层是使用**协程**(coroutine)实现并发，coroutine是一种运行在**用户态的用户线程**，类似于 greenthread

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

## 进程 vs 线程 vs 协程

对操作系统来说，**线程是最小的执行单元**，**进程是最小的资源管理单元**

![image](https://user-images.githubusercontent.com/13504729/130587050-cd540f9b-fdb9-4627-b550-a76e47cdae34.png)

协程，英文Coroutines，是一种比线程更加轻量级的存在。正如**一个进程可以拥有多个线程一样，一个线程也可以拥有多个协程**。**协程的调度完全由用户控制**。协程拥有自己的寄存器上下文和栈。

**协程与线程主要区别是它将不再被内核调度，而是交给了程序自己而线程是将自己交给内核调度
**
# GO并发模型-CSP理论

普通的线程并发模型，就是像Java、C++、或者Python，**他们线程间通信都是通过共享内存的方式来进行的。非常典型的方式就是，在访问共享数据（例如数组、Map、或者某个结构体或对象）的时候，通过锁来访问**，因此，在很多时候，衍生出一种方便操作的数据结构，叫做“**线程安全的数据结构**”。例如Java提供的包”**java.util.concurrent**”中的数据结构。Go中也实现了传统的线程并发模型

Golang的并发模型基于**CSP理论**，Golang并发的口号是：不用通过共享内存来通信，而是通过通信来共享内存。

Golang用来支持并发的元素集：
+ goroutines
+ channels
+ select
+ sync package
	
其中goroutines，channels和select 对应于实现CSP理论，即通过通信来共享内存。这几乎能解决Golang并发的90%问题，另外的10%场景需要通过同步原语来解决，即sync包相关的结构。

# 通道Channel

**“不要通过共享内存来通信，而应该通过通信来共享内存” 这是一句风靡golang社区的经典语**

```go
a := make(chan int) 

data := <- a // read from channel a  
a <- data // write to channel a

```

**发送和接收默认是阻塞的**

一个通道发送和接收数据，默认是阻塞的。当一个数据被发送到通道时，在发送语句中被阻塞，直到另一个Goroutine从该通道读取数据。相对地，当从通道读取数据时，读取被阻塞，直到一个Goroutine将数据写入该通道。

官方的go编译器限制channel最多能容纳到65535个元素，尽管如此，我们也不应该传递体积过大的元素值，因为channel的数据从进入到流出会涉及到数据拷贝操作。如果元素体积过大，最好的方法还是使用传递指针来取代传递值。

**channel类型是可以带有方向的，假设T是一种类型；chan T是双向channel类型，编译器允许对双向channel同时进行发送和接收；chan<- T是只写channel类型，编译器只允许往channel里面发送数据；<-chan T是只读channel类型，编辑器只允许从channel里面接收数据。**

**双向类型的channel，可以被强制转换成只读channel或者是只写channel，但是反过来却不行**，只读和只写channel是不可以转换成双向channel的。

channel里面的value buffer的容量也就是channel的容量。**channel的容量为零表示这是一个阻塞型通道，非零表示缓冲型通道[非阻塞型通道]**。

# select

Go中的select和channel配合使用，通过select可以监听多个channel的I/O读写事件，当IO操作发生时，触发相应的动作。

+ 若一个通道发生读写，则去执行对应逻辑  
+ 若多个通道发生读写，则随机选择一个去执行
+ 若多个通道都没有发生读写
	+ 有default分支：则走default分支
	+ 没有：则阻塞
	


**通道channel是将goroutine的粘合剂，select语句是通道channel的粘合剂**


![image](https://user-images.githubusercontent.com/13504729/130595261-76d0f65c-62f9-4286-b293-d2c89dbd84e0.png)

代码

```
var c1, c2 <-chan interface{}
var c3 chan<- interface{}
select {
case <-c1:
	// Do something
case <-c2:
	// Do something
case c3 <- struct{}{}:
	// Do something
}
```

# 特性

## 并发编程，协程goroutine
## 垃圾回收

垃圾对象判定算法
* 引用计数法 缺点：循环引用
* 可达性分析

垃圾回收算法
* 标记清除算法 缺点：1循环两遍 2产生碎片
* 标记整理算法 折中方案 解决碎片和内存利用率低
* 复制算法 缺点：内存利用率低
* 分代回收算法，分场景选择合适的回收算法 组合起来

golang-三色标记清扫算法

Golang的垃圾回收（GC）算法使用的是无分代（对象没有代际之分）、不整理（回收过程中不对对象进行移动与整理）、并发（与用户代码并发执行）的**三色标记清扫算法**。

三色标记法将对象分为三类，并用不同的颜色相称：
* 白色对象（可能死亡）：未被回收器访问到的对象。在回收开始阶段，所有对象均为白色，当回收结束后，白色对象均不可达。
* 灰色对象（波面）：已被回收器访问到的对象，但回收器需要对其中的一个或多个指针进行扫描，因为他们可能还指向白色对象。
* 黑色对象（确定存活）：已被回收器访问到的对象，其中所有字段都已被扫描，黑色对象中任何一个指针都不可能直接指向白色对象。

标记过程如下：

（1）起初所有的对象都是白色的；
（2）从根对象出发扫描所有可达对象，标记为灰色，放入待处理队列；
（3）从待处理队列中取出灰色对象，将其引用的对象标记为灰色并放入待处理队列中，自身标记为黑色；
（4）重复步骤（3），直到待处理队列为空，此时白色对象即为不可达的“垃圾”，回收白色对象；

![image](https://user-images.githubusercontent.com/13504729/149907851-b369ffb4-a8e9-4c25-bb46-511dfa501451.png)


## 网络编程，grpc

由于golang诞生在互联网时代，因此它天生具备了去中心化、分布式等特性，具体表现之一就是提供了丰富便捷的网络编程接口，
* socket用net.Dial(基于tcp/udp，封装了传统的connect、listen、accept等接口)
* http用http.Get/Post()
* rpc用client.Call('class_name.method_name', args, &reply)

## 内置常用数据类型,map slice

内置了一个其他静态语言通常用库方式来支持的**字典类型（map）**，理由很简单：既然绝大多数开发者都需要用到这个类型，为什么还非要每个人都写一行import语句来包含一个库？

Go语言还新增了一个数据类型：**数组切片（Slice）**。我们可以将数组切片看作是一种**可动态增长的数组**，其功能与C++标准库中的vector类似，但有效地消除了反复写以下几行代码的工作量：

```
#include <vector>
#include<map>
#include<algorithm>
using namespace std;
```

## 函数多返回值

Go语言的多返回值功能让开发者既不用再只为了返回多个值而专门定义一个数据结构

## 异常处理 defer

Golang 没有结构化异常，使用 panic 抛出错误，recover 捕获错误。
异常的使用场景简单描述：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

```
// 举例除数为0
func main() {
    var a1 int = 2
    var b1 int = 0
    // 犹豫除数为0 报错终止
    res1 := divisionIntRecover(a1, b1)
    fmt.Println("res1",res1)
}

func divisionIntRecover(a int, b int) (ret int) {
// 先注释掉 defer 进程终止抛出异常
    //defer func() {
    //    if err := recover(); err != nil {
    //        // 打印异常，关闭资源，退出此函数
    //        fmt.Println(err)
    //        ret = -3
    //    }
    //}()

    return a / b
}
```

## 支持匿名函数

```
f := func(x,y,int) int{
  return x+y
}
```

## 简单的类型，不支持继承和重载

Go语言的类型定义非常接近C语言中的结构（struct），甚至直接沿用了struct关键字。为了保持简洁，Go语言没有直接沿袭C++和Java的传统去设计一个超级复杂的类型系统，**不支持继承和重载**，而只是支持最基本的**类型组合功能**。虽然看起来过于简洁，但Go语言依然能够实现C++和Java使用复杂的类型系统才能实现的功能。



