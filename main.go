package main

import (
	"fmt"
	"github.com/didadadida93/learn-go-1/pkg/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/ping", handler.PingHandler)
	http.HandleFunc("/characters", handler.CharacterHandler)

	fmt.Println("listening on port :3000")
	http.ListenAndServe(":3000", nil)
}
