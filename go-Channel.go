package main

import (
    "time"
    "math/rand"
    "fmt"
)

func longTimeRequest() <-chan int32 {
    // 无缓冲通道，读空或者写满的操作都会阻塞哦
    r := make(chan int32)

    go func() {
        time.Sleep(time.Second * 3) // 模拟一个工作负载
        num := rand.Int31n(100)
        fmt.Println(num) 
        r <- num
    }()

    return r
}

func sumSquares(a, b int32) int32 {
    return a*a + b*b
}

func main() {
    rand.Seed(time.Now().UnixNano())

    a, b := longTimeRequest(), longTimeRequest()
    // 读空 阻塞中， 当写入线程写入后 继续执行
    fmt.Println(sumSquares(<-a, <-b))
}