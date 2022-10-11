package module2

import "net/http"

func main() {
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe()
}
