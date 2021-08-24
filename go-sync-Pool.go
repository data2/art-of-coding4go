package main

import (
    "sync"
    "fmt"
)

func main() {
    
    pool := &sync.Pool{
        New: func() interface{}{
            return "str"
        },
        
    }    

    fmt.Println(pool.Get())
    pool.Put("str_new")
    fmt.Println(pool.Get())


    myPool := &sync.Pool{
        New: func() interface{} {
            fmt.Println("Creating new instance.")
            return struct{}{}
        },
    }

    myPool.Get()             //1
    instance := myPool.Get() //1
    fmt.Println("-----分割----")
    myPool.Put(instance)     //2
    myPool.Get() 

    

}