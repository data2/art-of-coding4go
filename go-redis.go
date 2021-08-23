package main 

import (
  "fmt"
  "github.com/go-redis/redis"
)

func main(){
  fmt.Println("golang start conn redis ")
  
  client := redis.NewClient(&redis.Options{
        Addr: "192.168.8.200:6379",
        Password: "123456",
        DB: 0,
    })
  
  pong, err := client.Ping().Result()
  fmt.Println(pong, err)
}
