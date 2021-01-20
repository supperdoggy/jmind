package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// getBlockByNumber - send req to api and returns GetBlockByNumberStruct
func getBlockByNumber(block int64) (answer GetBlockByNumberStruct, err error) {
	// making req to api
	path := getApiLink(fmt.Sprintf("%x", block))
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("getBlockByNumber() error:", err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("getBlockByNumber() error reading body:", err.Error())
		return
	}
	defer resp.Body.Close()

	// resp we got from api
	if err := json.Unmarshal(body, &answer); err != nil {
		fmt.Println("error unmarshal", err.Error(), string(body))
	}
	return
}
