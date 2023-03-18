package handler

import (
	"encoding/json"
	"forevermzm/blockchain/pkg/core"
	"net/http"
)

var Blockchain core.Blockchain

func HandleBlockchain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		handleGetBlockchain(w, req)
	case "POST":
		handlePostBlockchain(w, req)
	}
}

func handleGetBlockchain(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(Blockchain)
}

func handlePostBlockchain(w http.ResponseWriter, req *http.Request) {
	var p core.Person
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	block := core.CreateBlock(Blockchain.Chain[len(Blockchain.Chain)-1], p)
	Blockchain, err = Blockchain.AddBlock(block)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
