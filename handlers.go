package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// getValueReq - returns to user number of transaction in given block and total amount of Ether
func getValueReq(writer http.ResponseWriter, request *http.Request) {
	// path and method validation
	if request.Method != "GET" {
		SendJsonAnswer(writer, http.StatusNotFound, []byte{})
		return
	}
	url := strings.Split(request.URL.Path, "/")
	if len(url) != 5 || url[4] != "total" {
		SendJsonAnswer(writer, http.StatusNotFound, []byte{})
		return
	}
	// getting and validating block id
	block := url[3]
	if block == "" || !regexp.MustCompile(`^[0-9]+`).MatchString(block) {
		SendJsonAnswer(writer, http.StatusNotFound, []byte{})
		return
	}

	// i didnt find anything smarter, sorry
	blocknum, err := strconv.ParseInt(block, 10, 64)
	if err != nil {
		SendJsonAnswer(writer, http.StatusBadRequest, []byte{})
		return
	}

	// check if we have value in cache, if so, return it to user
	if data, ok := cache.GetValue(fmt.Sprintf("%x", blocknum)); ok {
		fmt.Println("got from cache")
		answerJson, err := json.Marshal(data)
		if err != nil {
			SendJsonAnswer(writer, http.StatusBadRequest, []byte{})
			fmt.Println("Marshal error:", err.Error())
			return
		}
		SendJsonAnswer(writer, http.StatusOK, answerJson)
		return
	}

	// making req to api
	path := getApiLink(fmt.Sprintf("%x", blocknum))
	resp, err := http.Get(path)
	if err != nil {
		SendJsonAnswer(writer, http.StatusBadRequest, []byte{})
		fmt.Println("error making req", err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		SendJsonAnswer(writer, http.StatusBadRequest, []byte{})
		fmt.Println("error reading body", err.Error())
		return
	}
	defer resp.Body.Close()

	// resp we got from api
	var Answer struct {
		R Result `json:"result"`
	}

	if err := json.Unmarshal(body, &Answer); err != nil {
		data := obj{"error": "api key error, try again"}
		marshaled, _ := json.Marshal(data)
		SendJsonAnswer(writer, http.StatusBadRequest, marshaled)
		fmt.Println("error unmarshal", err.Error(), string(body))
		return
	}
	// enumerating total number of Ether
	amount := countValue(Answer.R.Transactions)
	// prepating answer
	answer := getValueReqAnswer{
		Amount:       amount,
		Transactions: len(Answer.R.Transactions),
	}
	answerJson, err := json.Marshal(answer)
	if err != nil {
		SendJsonAnswer(writer, http.StatusBadRequest, []byte{})
		fmt.Println("Marshal error:", err.Error())
		return
	}
	SendJsonAnswer(writer, http.StatusOK, answerJson)
	cache.WriteToCache(fmt.Sprintf("%x", blocknum), answer)
}
