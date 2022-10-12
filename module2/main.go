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
	w.WriteHeader(233)
	fmt.Println(r.Header)
	//for k := range r.Header {
	//	w.Header().Set(k, "value of this key")
	//}
	// w.Header().Set("reqheader", string(r.Header))
	io.WriteString(w, "ok")
}
