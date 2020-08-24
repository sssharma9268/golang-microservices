package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
	fmt.Println("hello")
}
