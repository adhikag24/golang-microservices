package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("hello world!"))
	})
	fmt.Println("Listening on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
