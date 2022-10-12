package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Println(k, v)
		for _, headerv := range v {
			w.Header().Add(k, headerv)
		}
	}
	//w.Header().Set()
	w.WriteHeader(233)
	//for k := range r.Header {
	//	w.Header().Set(k, "value of this key")
	//}
	// w.Header().Set("reqheader", string(r.Header))
	io.WriteString(w, "ok")
}
