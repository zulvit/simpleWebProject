package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", pageHandler)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
