package main

type Transaction struct {
	Value string `json:"value"`
}

type Result struct {
	Transactions []Transaction `json:"transactions"`
}
