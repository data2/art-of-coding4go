package main 

import (
  "fmt"
  "github.com/go-redis/redis"
)

func TestRedis(){
  fmt.Println("golang start conn redis ")
  
  client := redis.NewClient(&redis.Options{
        Addr: "20.26.52.89:6379",
        //Password: "123456",
        DB: 0,
    })
  
  pong, err := client.Ping().Result()
  fmt.Println(pong, err)

  err = client.Set("testKey","testGo",0).Err()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("go redis setKey success")

  value, err := client.Get("testKey").Result()
  fmt.Println(value)

}
