package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Result string `json:"result"`
}

func uppercaseHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req Request
	json.NewDecoder(r.Body).Decode(&req)

	upper := strings.ToUpper(req.Message)

	resp := Response{Result: upper}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/uppercase", uppercaseHandler)
	http.ListenAndServe(":8080", nil)
}
