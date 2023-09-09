// a go program to build a Simple Webserver
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})
	http.ListenAndServe(":8080", nil)
}
