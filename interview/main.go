package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Transaction struct {
	Amount    string
	Timestamp time.Time
}

type Response struct {
	Sum   float64
	Avg   float64
	Max   float64
	Min   float64
	Count int
}

var transactionCount int = 0
var transactionMap = make(map[int]Transaction,0)
var responseMap = make(map[int]Response,0)
var counterTime = time.Now().UTC()

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/transactions", CreateTransaction).Methods("POST")
	router.HandleFunc("/statistics", GetTransaction).Methods("GET")
	router.HandleFunc("/transactions", DeleteTransaction).Methods("DELETE")
	log.Println("Server connected at: 8000")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Panic(err.Error())
	}
}


func CreateTransaction(w http.ResponseWriter, r *http.Request){
	var t Transaction
	now := time.Now().UTC()

	res := responseMap[0]
	res.Count = transactionCount

	if now.After(counterTime){
		counterTime = now
		counterTime = now.Add(time.Second * 60)
		transactionMap = make(map[int]Transaction,0)
		responseMap = make(map[int]Response,0)
		transactionCount = 0
	}

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid JSON Format", http.StatusBadRequest)
		return
	}
	transactionMap[transactionCount] = t
	f, err := strconv.ParseFloat(transactionMap[transactionCount].Amount, 8)
	res.Sum += f
	if res.Max < f {
		res.Max = f
	}

	// var test2 float64 = 0

	if res.Min == 0 && transactionCount == 0 {
		res.Min = f
	}else if res.Min > f {
		res.Min = f
	}

	// Given input date is in UTC so fetching current time from UTC
	test := now.Add(-time.Second * 60)

	if now.After(transactionMap[transactionCount].Timestamp) { 
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	if test.After(transactionMap[transactionCount].Timestamp) { 
		w.WriteHeader(http.StatusNoContent)
		return
	}
	transactionCount = transactionCount + 1
	responseMap[0] = res

	w.WriteHeader(http.StatusCreated)
	returnStatement := "Data Inserted for " + fmt.Sprintf("%d",transactionCount)
	w.Write([]byte(returnStatement))
}

func GetTransaction(w http.ResponseWriter, r *http.Request){
	now := time.Now().UTC()

	if now.After(counterTime){
		counterTime = now
		counterTime = now.Add(time.Second * 60)
		transactionMap = make(map[int]Transaction,0)
		responseMap = make(map[int]Response,0)
		transactionCount = 0
	}
	var res Response

	res.Sum = responseMap[0].Sum
	res.Avg = res.Sum / float64(transactionCount)
	res.Max = responseMap[0].Max
	res.Min = responseMap[0].Min
	res.Count = transactionCount

	json.NewEncoder(w).Encode(res)

}

func DeleteTransaction(w http.ResponseWriter, r *http.Request){
	transactionMap = make(map[int]Transaction,0)
	responseMap = make(map[int]Response,0)
	transactionCount = 0
	w.WriteHeader(http.StatusNoContent)

	w.Write([]byte("Data deleted"))
}
