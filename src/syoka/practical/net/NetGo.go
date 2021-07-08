package main

import (
	"net/http"
	"syoka/practical/utils"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write(utils.ToByte("hello boy"))
}

func main() {
	http.HandleFunc("/", sayHello)
	http.ListenAndServe("localhost:8080", nil)
}
