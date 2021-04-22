package main

import (
	"fmt"
	"github.com/didadadida93/learn-go-1/pkg"
	"net/http"
)

func main() {
	http.HandleFunc("/", pkg.RootHandler)
	http.HandleFunc("/ping", pkg.PingHandler)
	http.HandleFunc("/characters", pkg.CharacterHandler)

	fmt.Println("listening on port :3000")
	http.ListenAndServe(":3000", nil)
}
