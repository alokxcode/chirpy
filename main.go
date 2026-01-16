package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.NewServeMux()

	handler.Handle("/", http.FileServer(http.Dir(".")))

	server := http.Server{
		Addr:    ":8081",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
