package main

import (
    "fmt"
    "sync"
)

var (
    dataMap = make(map[int]int)
    mutex   = &sync.Mutex{}
)

func writeToMap(key int, value int) {
    mutex.Lock()
    defer mutex.Unlock()
    dataMap[key] = value
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            writeToMap(i, i)
        }(i)
    }

    wg.Wait()

    for key, value := range dataMap {
        fmt.Printf("Key: %d, Value: %d\n", key, value)
    }
}

