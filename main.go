package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	cache = Cache{
		M:   map[string]getValueReqAnswer{},
		mut: sync.Mutex{},
	}
	ApiKey = getApiKey()
)

func main() {
	fmt.Println("Started server...")

	http.HandleFunc("/api/block/", getValueReq)
	// staring server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err.Error())
	}
}
