package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/m/", debugHandler)
	mux.HandleFunc("/m/work/simplestamp", debugHandler)
	fmt.Println("listen:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func debugHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	q, _ := url.ParseQuery(string(b))
	fmt.Println(r.Method, r.URL, q)
	fmt.Fprintf(w, "ok")
}
