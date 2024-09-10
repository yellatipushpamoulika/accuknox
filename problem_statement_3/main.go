package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Current Date and Time: %s", time.Now().Format(time.RFC1123))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9090", nil)
}
