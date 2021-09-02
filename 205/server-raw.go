package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"Code":0,"Message":"Success","RequestID":"1234-5678-90","PublisherList":[]}`)
}

func main() {
	http.HandleFunc("/decoder/publisher", IndexHandler)
	http.ListenAndServe(":8000", nil)
}
