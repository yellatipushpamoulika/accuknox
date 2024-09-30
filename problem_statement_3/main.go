package main

import (
    "fmt"
    "time"
)

func main() {
    curtime:=time.Now()
    formattime:=curtime.Format("2006-01-02 15:04:05")
    fmt.Println(formattime)
}
