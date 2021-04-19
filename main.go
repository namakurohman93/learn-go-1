package main

import (
    "fmt"
    "net/http"
    "github.com/didadadida93/go-learn-1/pkg/handler"
)

func main() {
    http.HandleFunc("/", handler.RootHandler)
    http.HandleFunc("/ping", handler.PingHandler)

    fmt.Println("listening on port :3000")
    http.ListenAndServe(":3000", nil)
}
