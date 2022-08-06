package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"Code":0,"Message":"Success"}`)
}

func main() {
	http.HandleFunc("/httpraw/test", IndexHandler)
	http.ListenAndServe(":8000", nil)
}
