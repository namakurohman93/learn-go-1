package handler

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("root handler\n"))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("ping handler\n"))
}
