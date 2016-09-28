package main

import (
    "fmt"
    "time"
)


func bigBytes() *[]byte {
        s := make([]byte, 100000000)
        return &s
}

func main() {
    fmt.Println("Starting DC/OS-101  OOM App ")
    var store map[int]*[]byte
    store = make(map[int]*[]byte)
    for i := 0; i < 500; i++ {
      store[i] =  bigBytes()
      fmt.Println("Eating Memory")
      time.Sleep(100 * time.Millisecond)
   }
}
