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
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	url := strings.Split(request.URL.Path, "/")
	if len(url) != 5 || url[4] != "total" {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	// getting and validating block id
	block := url[3]
	if block == "" || !regexp.MustCompile(`^[0-9]+`).MatchString(block) {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// i didnt find anything smarter, sorry
	blocknum, err := strconv.ParseInt(block, 10, 64)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// making req to api
	path := getApiLink(fmt.Sprintf("%x", blocknum))
	resp, err := http.Get(path)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println("error making req", err.Error())
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println("error reading body", err.Error())
		return
	}
	defer resp.Body.Close()

	// resp we got from api
	var Answer struct {
		R Result `json:"result"`
	}

	if err := json.Unmarshal(body, &Answer); err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println("error unmarshal", err.Error())
		return
	}
	// enumerating total number of Ether
	amount := countValue(Answer.R.Transactions)
	// prepating answer
	answer := obj{"amount": amount, "transactions": len(Answer.R.Transactions)}

	answerJson, err := json.Marshal(answer)
	if err != nil {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Println("Marshal error:", err.Error())
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(answerJson)
}
