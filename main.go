package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/api/block/", getValueReq)
	// staring server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}
