package main

import "net/http"

func statusPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Status Unknown"))
}

func main() {
	http.HandleFunc("/", statusPage)
	http.ListenAndServe(":8080", nil)
}
