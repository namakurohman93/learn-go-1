package handler

import (
	"github.com/didadadida93/learn-go-1/pkg/character"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root handler\n"))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping handler\n"))
}

func CharacterHandler(w http.ResponseWriter, r *http.Request) {
	characterByte, err := character.GetCharacterByte()
	if err != nil {
		w.Write([]byte("error has been occured"))
		return
	}

	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.Write(characterByte)
}
