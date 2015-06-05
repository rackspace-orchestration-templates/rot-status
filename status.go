package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", statusPage)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on :8080 ...")
}

func someFunc(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Status Unknown"))
}
