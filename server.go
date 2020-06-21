package main

import (
	"fmt"
	"net/http"
)

const port = ":8090"

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("."))))
	fmt.Println("listening on port " + port)
	http.ListenAndServe(port, nil)
}
