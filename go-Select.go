package main

import (
    "fmt"
    "time"
)

func main() {
    c1 := make(chan string, 1)
    c2 := make(chan string, 1)

    // 开启一个协程，可以发送数据到信道
    go func() {
        time.Sleep(time.Second * 1)
        c2 <- "hello"
    }()

    select {
      case msg1 := <-c1:
        fmt.Println("c1 received: ", msg1)
      case msg2 := <-c2:
        fmt.Println("c2 received: ", msg2)
      // default:
      //   fmt.Println("No data received.")
    }

    fmt.Println("结束")
}