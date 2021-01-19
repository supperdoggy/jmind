package main

import (
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

type obj map[string]interface{}

const (
	Ether Factor = 18
)

type Factor int

// Unit Base struct
type Unit struct {
	Name      string
	WeiFactor *big.Int
}

// GetWeiFactor func return unit.WeiFactor
func (unit *Unit) GetWeiFactor() *big.Int {
	return unit.WeiFactor
}

// InitUnit func param unitname(string) retun Unit struct
// by default Wei factor
func (unit *Unit) InitUnit() *Unit {
	units := Ether
	var unitfactor, exp = big.NewInt(int64(units)), big.NewInt(10)
	exp.Exp(exp, unitfactor, nil)
	unit.Name = "ether"
	unit.WeiFactor = exp

	return unit
}

// FromWei func calculate number / unitfactor return string
func FromWei(number string) string {
	unit := new(Unit)
	unit = unit.InitUnit()
	bigFloatNumber, _ := new(big.Float).SetString(number)
	unitWeiFactor := new(big.Float).SetInt(unit.GetWeiFactor())
	tmp := new(big.Float)
	tmp.Quo(bigFloatNumber, unitWeiFactor)
	tmpstr := fmt.Sprintf("%v", tmp)

	return tmpstr
}

// intToHexa - transforms decimal to hexadecimal
func intToHexa(num int) string {
	return strconv.FormatInt(int64(num), 16)
}

// countValue - returns total number of Ether in string
func countValue(transactions []Transaction) string {
	var total float64
	for _, v := range transactions {
		numStr := FromWei(v.Value)
		num, _ := strconv.ParseFloat(numStr, 64)
		total += num
	}
	return fmt.Sprintf("%.6f", total)
}

// getApiLink- generates api link
func getApiLink(tag string) string {
	return fmt.Sprintf("https://api.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%v&boolean=true&apikey=YourApiKeyToken", tag)
}

// SendJsonAnswer - sends answer with given data and status code
func SendJsonAnswer(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}
