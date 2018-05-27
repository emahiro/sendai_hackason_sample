package main

import (
	"fmt"
	"net/http"

	"sendaihackathonsample/handler"

	"github.com/gorilla/mux"
)

var port  = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/article", handler.Index)

	if err := http.ListenAndServe(port, r); err != nil{
		fmt.Errorf("[ERROR] Cannot start web server")
		return
	}
}