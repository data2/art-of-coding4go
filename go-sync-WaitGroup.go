package main

import (
    "fmt"
    "sync"
)


// go run go-sync-WaitGroup.go  go-http.go
func main(){
    fmt.Println("10个线程开始执行")

	var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go func(){
            defer wg.Done()
            TestHttp()
        }()
    }

    wg.Wait()
    fmt.Println("10个线程执行完毕")
}