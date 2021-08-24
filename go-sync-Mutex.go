package main

import (
    "fmt"
    "sync"
)

var lock sync.Mutex
var wg sync.WaitGroup
var count int

func increase(){
        lock.Lock()
        defer lock.Unlock()
        count++
        fmt.Printf("Incrementing: %d\n", count)
    
}

func decrease(){
        lock.Lock()
        defer lock.Unlock()
        count--
        fmt.Printf("Incrementing: %d\n", count)
}

func main(){
    for i := 0 ; i<5; i++ {
        wg.Add(1)
        go func(){
            defer wg.Done()
            increase()
        }()
    }

    for i := 0 ; i<5; i++ {
        wg.Add(1)
        go func(){
            defer wg.Done()
            decrease()
        }()
    }

    wg.Wait()
}