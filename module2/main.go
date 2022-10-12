package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	//get request header and put them into response header
	for k, v := range r.Header {
		//fmt.Println(k, v)
		for _, headerv := range v {
			w.Header().Add(k, headerv)
		}
	}
	//get os version and put it into response header
	version := os.Getenv("VERSION")
	fmt.Printf("VERSION = %s\n", version)
	if version == "" {
		version = "Not Available!"
	}
	w.Header().Add("VERSION", version)

	//set response statuscode
	w.WriteHeader(233)

	//se response content
	io.WriteString(w, "ok")
}
