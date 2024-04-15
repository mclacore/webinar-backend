package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type FizzBuzzResult struct {
	Result []string `json:"result"`
}

func fizzBuzz(n int) []string {
	results := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			results[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			results[i-1] = "Fizz"
		} else if i%5 == 0 {
			results[i-1] = "Buzz"
		} else {
			results[i-1] = strconv.Itoa(i)
		}
	}
	return results
}

func fizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		http.Error(w, "Missing 'n' parameter", http.StatusBadRequest)
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil {
		http.Error(w, "'n' must be an integer", http.StatusBadRequest)
		return
	}

	if n < 1 {
		http.Error(w, "'n' must be greater than 0", http.StatusBadRequest)
		return
	}

	results := fizzBuzz(n)
	result := FizzBuzzResult{Result: results}
	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error processing results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/fizzbuzz", fizzBuzzHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
