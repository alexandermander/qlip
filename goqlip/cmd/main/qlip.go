package main

// i want to host a api server that will be able to handle requests from the client
import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	//start the server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":443", nil)
}

