package main

import (
	"fmt"
	"sync"
)

type Transaction struct {
	Value string `json:"value"`
}

type Result struct {
	Transactions []Transaction `json:"transactions"`
}

type GetBlockByNumberStruct struct {
	R Result `json:"result"`
}

type getValueReqAnswer struct {
	Amount       string `json:"amount"`
	Transactions int    `json:"transactions"`
}

// Cache - cache for getValueReqAnswer
type Cache struct {
	M   map[string]getValueReqAnswer
	mut sync.Mutex
}

// GetValue - Returns cached value
func (c *Cache) GetValue(id string) (answer getValueReqAnswer, ok bool) {
	c.mut.Lock()
	answer, ok = c.M[id]
	c.mut.Unlock()
	return
}

// WriteToCache - writes value to cache
func (c *Cache) WriteToCache(id string, data getValueReqAnswer) {
	c.mut.Lock()
	c.M[id] = data
	c.mut.Unlock()
	fmt.Println("writing to cache:", id)
}
